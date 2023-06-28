package orbit

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

/****************************************** Basic Interface ************************************************/

type Satellite interface {
	Meta(context.Context) (*proto.Metadata, error)
	Symbols(context.Context) ([]vm.FunctionDefinition, error)
	Call(ctx context.Context, module vm.Module, function string, inputs []uint64) ([]uint64, error)
}

type Module interface {
	MemoryRead(offset uint32, size uint32) ([]byte, error)
	MemoryWrite(offset uint32, data []byte) (uint32, error)
}

var _ vm.FunctionDefinition = &functionDefinition{}

type functionDefinition struct {
	name string
	args []vm.ValueType
	rets []vm.ValueType
}

func (f *functionDefinition) Name() string {
	return f.name
}

func (f *functionDefinition) ParamTypes() []vm.ValueType {
	return f.args
}

func (f *functionDefinition) ResultTypes() []vm.ValueType {
	return f.rets
}

/****************************************** PLUGIN IFACE *******************************/

type satellite struct {
	plugin.NetRPCUnsupportedPlugin

	name    string
	exports map[string]interface{}
}

type link struct {
	plugin.NetRPCUnsupportedPlugin
}

var _ Satellite = &GRPCPluginClient{}

func NewSatellite(
	// turn this shit to an interface?
	name string,
	exports func() map[string]interface{},
) plugin.Plugin {
	return &satellite{
		name:    name,
		exports: exports(),
	}
}

func NewLink() plugin.Plugin {
	return &link{}
}

type ModulePlugin struct {
	plugin.Plugin

	Module vm.Module
}

func (st *satellite) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, &GRPCPluginServer{
		broker:    broker,
		satellite: st,
	})

	return nil
}

func (p *satellite) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("can't create a link (satellite client) from main process")

}

func (p *link) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("can't create a satelite server from main process")
}

func (p *link) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCPluginClient{
		client: proto.NewPluginClient(c),
		broker: broker,
	}, nil
}

// var ctxType = reflect.TypeOf((context.Background())).Elem()
var moduleType = reflect.TypeOf((*Module)(nil)).Elem()

func (p *GRPCPluginServer) Symbols(context.Context, *proto.Empty) (*proto.FunctionDefinitions, error) {
	ret := &proto.FunctionDefinitions{
		Functions: make([]*proto.FunctionDefinition, 0, len(p.satellite.exports)),
	}
	for name, handler := range p.satellite.exports {
		fx := reflect.TypeOf(handler)
		if fx.Kind() != reflect.Func {
			return nil, fmt.Errorf("handler %s for not a function", name)
		}

		argsType := make([]proto.Type, 0, fx.NumIn())
		for i := 0; i < fx.NumIn(); i++ {
			if (i == 0 && fx.In(i).Implements(ctxType)) || (i == 1 && fx.In(i).Implements(moduleType)) {
				continue
			}
			switch fx.In(i).Kind() {
			case reflect.Int32, reflect.Uint32:
				argsType = append(argsType, proto.Type_i32)
			case reflect.Int64, reflect.Uint64:
				argsType = append(argsType, proto.Type_i64)
			case reflect.Float32:
				argsType = append(argsType, proto.Type_f32)
			case reflect.Float64:
				argsType = append(argsType, proto.Type_f64)
			}
		}

		retTypes := make([]proto.Type, 0, fx.NumOut())
		for i := 0; i < fx.NumOut(); i++ {
			switch fx.Out(i).Kind() {
			case reflect.Int32, reflect.Uint32:
				retTypes = append(retTypes, proto.Type_i32)
			case reflect.Int64, reflect.Uint64:
				retTypes = append(retTypes, proto.Type_i64)
			case reflect.Float32:
				retTypes = append(retTypes, proto.Type_f32)
			case reflect.Float64:
				retTypes = append(retTypes, proto.Type_f64)
			}
		}

		ret.Functions = append(ret.Functions, &proto.FunctionDefinition{
			Name: name,
			Args: argsType,
			Rets: retTypes,
		})
	}

	return ret, nil
}

func (p *GRPCPluginServer) Meta(context.Context, *proto.Empty) (*proto.Metadata, error) {
	return &proto.Metadata{
		Name: p.satellite.name,
	}, nil
}

func (p *GRPCPluginServer) Call(ctx context.Context, req *proto.CallRequest) (*proto.CallReturn, error) {
	f, _ := os.Create("/tmp/log.txt")
	defer f.Close()
	fmt.Fprintf(f, ">>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n")
	time.Sleep(time.Second)
	conn, err := p.broker.Dial(req.Broker)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	mod := NewModuleLink(ctx, conn)

	handler, ok := p.satellite.exports[req.Function]
	if !ok {
		return nil, errors.New("bitch")
	}

	fx := reflect.ValueOf(handler)
	tfx := fx.Type()
	in := make([]reflect.Value, 0, len(req.Inputs)+2)

	if tfx.NumIn() >= 1 && tfx.In(0) == ctxType {
		in = append(in, reflect.ValueOf(ctx))
	}

	if tfx.NumIn() >= 2 && tfx.In(1) == moduleType {
		in = append(in, reflect.ValueOf(mod))
	}

	for _, v := range req.Inputs {
		var rv reflect.Value
		switch tfx.In(len(in)).Kind() {
		case reflect.Int32:
			rv = reflect.ValueOf(int32(v))
		case reflect.Int64:
			rv = reflect.ValueOf(int64(v))
		case reflect.Uint32:
			rv = reflect.ValueOf(uint32(v))
		case reflect.Uint64:
			rv = reflect.ValueOf(uint64(v))
		default:
			return nil, fmt.Errorf("---------------")
		}
		in = append(in, rv)
	}

	fmt.Fprintf(f, ">>>>>>>>>>%#v\n", in)
	out := fx.Call(in)

	ret := make([]uint64, len(out))
	for i, _arg := range out {
		switch _arg.Kind() {
		case reflect.Float32:
			ret[i] = uint64(math.Float32bits(float32(_arg.Float())))
		case reflect.Float64:
			ret[i] = math.Float64bits(_arg.Float())
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			ret[i] = _arg.Uint()
		case reflect.Int, reflect.Int32, reflect.Int64:
			ret[i] = uint64(_arg.Int())
		default:
			return nil, fmt.Errorf("failed to process arguments %v of type %T", _arg, _arg)
		}
	}

	return &proto.CallReturn{
		Rets: ret,
	}, nil
}

/************************************************************************/

type moduleLink struct {
	plugin.NetRPCUnsupportedPlugin
	ctx    context.Context
	client proto.ModuleClient
}

type module struct {
	proto.UnimplementedModuleServer
	module vm.Module
}

func NewModuleLink(ctx context.Context, conn *grpc.ClientConn) Module {
	return &moduleLink{ctx: ctx, client: proto.NewModuleClient(conn)}
}

func NewModule(mod vm.Module) proto.ModuleServer {
	return &module{module: mod}
}

func (p *moduleLink) MemoryRead(offset uint32, size uint32) ([]byte, error) {
	ret, err := p.client.MemoryRead(p.ctx, &proto.ReadRequest{Offset: offset, Size: size})
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (p *moduleLink) MemoryWrite(offset uint32, data []byte) (uint32, error) {
	ret, err := p.client.MemoryWrite(p.ctx, &proto.WriteRequest{Offset: offset, Data: data})
	if err != nil {
		return 0, err
	}
	if ret.Error != proto.IOError_none && ret.Error != proto.IOError_eof {
		return 0, ret.Error.Error()
	}
	return uint32(len(data)), ret.Error.Error()
}

// var _ Plugin = &GRPCPluginClient{}

// func (p *ModulePlugin) MemoryRead(offset uint32, size uint32) ([]byte, error) {
// 	data, ok := p.Module.Memory().Read(offset, size)
// 	if !ok {
// 		return nil, io.EOF
// 	}

// 	return data, nil
// }

// func (p *ModulePlugin) MemoryWrite(offset uint32, data []byte) (uint32, error) {
// 	if ok := p.Module.Memory().Write(offset, data); !ok {
// 		return 0, io.ErrShortWrite
// 	}

// 	return uint32(len(data)), nil
// }

// func (p *moduleLink) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
// 	return errors.New("make me bitch")
// }

// func (p *moduleLink) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
// 	_ctx, _ctxC := context.WithCancel(ctx)
// 	return &GRPCModuleClient{
// 		client: proto.NewModuleClient(c),
// 		broker: broker,
// 		ctx:    _ctx,
// 		ctxC:   _ctxC,
// 	}, nil
// }

func (m *module) MemoryRead(ctx context.Context, req *proto.ReadRequest) (*proto.ReadReturn, error) {
	data, ok := m.module.Memory().Read(req.Offset, req.Size)
	if !ok {
		return nil, errors.New("bitch")
	}

	return &proto.ReadReturn{Data: data}, nil
}

func (m *module) MemoryWrite(ctx context.Context, req *proto.WriteRequest) (*proto.WriteReturn, error) {
	ok := m.module.Memory().Write(req.Offset, req.Data)
	if !ok {
		return nil, errors.New("bitch")
	}

	return &proto.WriteReturn{Written: uint32(len(req.Data))}, nil
}

/******************************************************/

func (c *GRPCPluginClient) Symbols(ctx context.Context) ([]vm.FunctionDefinition, error) {
	resp, err := c.client.Symbols(ctx, &proto.Empty{})
	if err != nil {
		return nil, err
	}

	funcDefs := make([]vm.FunctionDefinition, len(resp.Functions))
	for idx, function := range resp.Functions {
		args, err := typesToBytes(function.Args)
		if err != nil {
			return nil, fmt.Errorf("getting args failed with: %s", err)
		}

		rets, err := typesToBytes(function.Rets)
		if err != nil {
			return nil, fmt.Errorf("getting returns failed with: %s", err)
		}

		funcDefs[idx] = &functionDefinition{
			name: function.Name,
			args: args,
			rets: rets,
		}
	}

	return funcDefs, nil
}

func (c *GRPCPluginClient) Meta(ctx context.Context) (*proto.Metadata, error) {
	return c.client.Meta(ctx, &proto.Empty{})
}

func (c *GRPCPluginClient) Call(ctx context.Context, module vm.Module, function string, inputs []uint64) ([]uint64, error) {
	moduleServer := NewModule(module)

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)

		proto.RegisterModuleServer(s, moduleServer)

		return s
	}

	brokerID := c.broker.NextId()
	go c.broker.AcceptAndServe(brokerID, serverFunc)

	resp, err := c.client.Call(ctx, &proto.CallRequest{
		Broker:   brokerID,
		Function: function,
		Inputs:   inputs,
	})
	if err != nil {
		return nil, err
	}
	defer s.Stop()

	if resp.Error != nil {
		if resp.Error.Code != nil {
			return nil, fmt.Errorf("failed with code: %d", *resp.Error.Code)
		}

		return nil, fmt.Errorf("failed with: %s", resp.Error.Message)
	}

	return resp.Rets, nil
}

/*************************************** CLIENT/ SERVER **********************************/

// this is used by main process to communicate with plugin process
type GRPCPluginClient struct {
	broker *plugin.GRPCBroker
	client proto.PluginClient
}

// this is used by the plugin process to provide GRPC for GRPCPluginClient (main process)
type GRPCPluginServer struct {
	broker *plugin.GRPCBroker
	proto.UnimplementedPluginServer

	satellite *satellite
}

// type GRPCModuleClient struct {
// 	client proto.ModuleClient
// 	broker *plugin.GRPCBroker
// 	ctx    context.Context
// 	ctxC   context.CancelFunc
// }

// type GRPCModuleServer struct {
// 	module vm.Module

// 	broker *plugin.GRPCBroker
// 	proto.UnimplementedModuleServer
// }
