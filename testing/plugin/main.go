package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	orbit "github.com/taubyte/vm-orbit"
)

// type testPlugin struct{}

// type testFuncDef struct{}

// func (*testFuncDef) Name() string {
// 	return "upperCase"
// }

// func (*testFuncDef) ParamTypes() []vm.ValueType {
// 	// inPtr ,inLen, outPtr
// 	return []vm.ValueType{vm.ValueTypeExternref, vm.ValueTypeExternref, vm.ValueTypeI64, vm.ValueTypeI64, vm.ValueTypeI64}
// }
// func (*testFuncDef) ResultTypes() []vm.ValueType {
// 	return []vm.ValueType{vm.ValueTypeI64}
// }

// func (*testPlugin) Symbols() ([]vm.FunctionDefinition, error) {
// 	return []vm.FunctionDefinition{&testFuncDef{}}, nil
// }

// func (*testPlugin) Call(ctx context.Context, module orbit.Module, function string, inputs []uint64) ([]uint64, error) {
// 	if len(inputs) != 5 {
// 		return nil, errors.New("expected 5 inputs")
// 	}

// 	inPtr, inLen, outPtr := inputs[2], inputs[3], inputs[4]

// 	data, err := module.MemoryRead(uint32(inPtr), uint32(inLen))
// 	if err != nil {
// 		return nil, err
// 	}

// 	upper := strings.ToUpper(string(data))
// 	_, err = module.MemoryWrite(uint32(outPtr), []byte(upper))
// 	if err != nil {
// 		return nil, err
// 	}

//		return []uint64{0}, nil
//	}
func (*testPlugin) Name() string {
	return "testplugin"
}

func hello(ctx context.Context, module orbit.Module) {
	fmt.Println("nlkjkljl")
}

func sum(a, b int64) int64 {
	return a + b
}

func exports() map[string]interface{} {
	return map[string]interface{}{
		"hello": hello,
		"sum":   sum,
	}
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: orbit.Handshake,
		Plugins: map[string]plugin.Plugin{
			"satellite": orbit.NewSatellite(
				"aladdin",
				exports,
			),
		},
	})
}
