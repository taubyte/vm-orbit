package orbit

import (
	"context"
	"errors"
	"os/exec"
	"reflect"

	goPlugin "github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
)

type vmPlugin struct {
	client  *goPlugin.Client
	address string
	name    string
}

type pluginInstance struct {
	plugin   *vmPlugin
	instance vm.Instance
	iface    Satellite
}

func Load(ma, name string) (vm.Plugin, error) {
	// ma is multiaddress
	p := &vmPlugin{address: ma, name: name}
	p.client = goPlugin.NewClient(&goPlugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         ClientPluginMap,
		Cmd:             exec.Command(ma),
		AllowedProtocols: []goPlugin.Protocol{
			goPlugin.ProtocolGRPC,
		},
	})

	return p, nil
}

func (p *vmPlugin) Name() string {
	return p.name
}

func (p *vmPlugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	rpcClient, err := p.client.Client()
	if err != nil {
		return nil, err
	}

	go func() {
		select {
		case <-instance.Context().Context().Done():
			rpcClient.Close()
		}
	}()

	raw, err := rpcClient.Dispense("satellite")
	if err != nil {
		return nil, err
	}

	pluginClient, ok := raw.(Satellite)
	if !ok {
		return nil, errors.New("not plugin")
	}

	pI := &pluginInstance{
		iface:    pluginClient,
		plugin:   p,
		instance: instance,
	}

	return pI, nil
}

var ctxType = reflect.TypeOf((context.Background())).Elem()
var vmMmoduleType = reflect.TypeOf((vm.Module)(nil)).Elem()
var I32Type = reflect.TypeOf((int32)(0)).Elem()
var I64Type = reflect.TypeOf((int64)(0)).Elem()

func (p *pluginInstance) convertToHandler(def vm.FunctionDefinition) (interface{}, error) {
	in := def.ParamTypes()
	_in := make([]reflect.Type, len(in)+2)

	_in[0] = ctxType
	_in[1] = vmMmoduleType

	for idx, pt := range in {
		switch pt {
		case vm.ValueTypeI32:
			_in[idx+2] = I32Type
		case vm.ValueTypeI64:
			_in[idx+2] = I64Type
		}
	}

	out := def.ResultTypes()
	_out := make([]reflect.Type, len(out))

	for idx, pt := range out {
		switch pt {
		case vm.ValueTypeI32:
			_in[idx] = I32Type
		case vm.ValueTypeI64:
			_in[idx] = I64Type
		}
	}

	_func := reflect.MakeFunc(
		reflect.FuncOf(_in, _out, false),
		func(args []reflect.Value) []reflect.Value {
			if len(args) < 2 {
				panic("")
			}

			ctx, ok := args[0].Interface().(context.Context)
			if !ok {
				panic("")
			}
			module, ok := args[1].Interface().(Module)
			if !ok {
				panic("")
			}

			_in := make([]uint64, len(args)-2)
			for i := 2; i < len(args); i++ {
				// TODO: double check uint64(int64) makes just a type conversion
				_in[i] = uint64(args[i].Int())
			}

			cOut, err := p.iface.Call(ctx, module, def.Name(), _in)
			if err != nil {
				panic(err)
			}

			_out := make([]reflect.Value, len(cOut))
			for idx := 0; idx < len(cOut); idx++ {
				switch out[idx] {
				case vm.ValueTypeI32:
					_out[idx] = reflect.ValueOf(int32(cOut[idx]))
				case vm.ValueTypeI64:
					_out[idx] = reflect.ValueOf(int64(cOut[idx]))
				}
			}

			return _out
		})

	return _func.Interface(), nil
}

func (p *pluginInstance) Load(hm vm.HostModule) (vm.ModuleInstance, error) {
	defs, err := p.iface.Symbols()
	if err != nil {
		return nil, err
	}

	for _, def := range defs {
		h, err := p.convertToHandler(def)
		if err != nil {
			return nil, err
		}
		hm.Functions(&vm.HostModuleFunctionDefinition{
			Name:    def.Name(),
			Handler: h,
		})
	}
	return nil, nil
}

func (p *pluginInstance) Close() error {
	return nil
}

func (p *pluginInstance) LoadFactory(factory vm.Factory, hm vm.HostModule) error {
	return nil
}
