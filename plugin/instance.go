package plugin

import (
	"context"
	"fmt"
	"reflect"

	"github.com/taubyte/go-interfaces/vm"
)

func (p *pluginInstance) Load(hm vm.HostModule) (vm.ModuleInstance, error) {
	defs, err := p.satellite.Symbols(p.instance.Context().Context())
	if err != nil {
		return nil, fmt.Errorf("getting (satellite) symbols failed with: %w", err)
	}

	funcDefs := make([]*vm.HostModuleFunctionDefinition, len(defs))
	for idx, def := range defs {
		funcDefs[idx] = &vm.HostModuleFunctionDefinition{
			Name:    def.Name(),
			Handler: p.convertToHandler(def),
		}
	}

	hm.Functions(funcDefs...)
	return hm.Compile()
}

func (p *pluginInstance) convertToHandler(def vm.FunctionDefinition) interface{} {
	in := bytesToReflect(def.ParamTypes(), []reflect.Type{vm.ContextType, vm.ModuleType})
	out := bytesToReflect(def.ResultTypes(), nil)

	return p.makeFunc(in, out, def).Interface()
}

// TODO: Handle Floats
func (p *pluginInstance) makeFunc(paramTypes []reflect.Type, retTypes []reflect.Type, def vm.FunctionDefinition) reflect.Value {
	return reflect.MakeFunc(
		reflect.FuncOf(paramTypes, retTypes, false),
		func(args []reflect.Value) []reflect.Value {
			if len(args) < 2 {
				panic("invalid function argument count, expected minimum 2")
			}

			ctx, ok := args[0].Interface().(context.Context)
			if !ok {
				panic("expected first argument to be context")
			}
			module, ok := args[1].Interface().(vm.Module)
			if !ok {
				panic("expected second argument to be vm.Module")
			}

			in := make([]uint64, 0, len(args))
			for i := 2; i < len(args); i++ {
				in = append(in, uint64(args[i].Int()))
			}

			cOut, err := p.satellite.Call(ctx, module, def.Name(), in)
			if err != nil {
				panic(fmt.Sprintf("calling `%s/%s` failed with: %s", module, def.Name(), err))
			}

			_out := make([]reflect.Value, len(cOut))
			for idx := 0; idx < len(cOut); idx++ {
				switch retTypes[idx] {
				case vm.I32Type:
					_out[idx] = reflect.ValueOf(int32(cOut[idx]))
				case vm.I64Type:
					_out[idx] = reflect.ValueOf(int64(cOut[idx]))
				}
			}

			return _out
		})
}

func (p *pluginInstance) Close() error {
	return nil
}

func bytesToReflect(raw []vm.ValueType, defaults []reflect.Type) []reflect.Type {
	types := make([]reflect.Type, 0, len(defaults)+len(raw))
	types = append(types, defaults...)

	for _, rawType := range raw {
		switch rawType {
		case vm.ValueTypeI32, vm.ValueTypeI64, vm.ValueTypeF32, vm.ValueTypeF64:
			types = append(types, vm.ValueTypeToReflectType(rawType))
		}
	}

	return types
}
