package common

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

var _ vm.FunctionDefinition = &functionDefinition{}

type functionDefinition struct {
	name string
	args []vm.ValueType
	rets []vm.ValueType
}

type Satellite interface {
	Meta(context.Context) (*proto.Metadata, error)
	Symbols(context.Context) ([]vm.FunctionDefinition, error)
	Call(ctx context.Context, module vm.Module, function string, inputs []uint64) ([]uint64, error)
}

type Module interface {
	MemoryRead(offset uint32, size uint32) ([]byte, error)
	MemoryWrite(offset uint32, data []byte) (uint32, error)
}
