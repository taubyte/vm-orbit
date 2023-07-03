package vm

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

type pluginInstance struct {
	plugin    *vmPlugin
	instance  vm.Instance
	satellite Satellite
}

type vmPlugin struct {
	client  *plugin.Client
	address string
	name    string
}

type Satellite interface {
	Meta(context.Context) (*proto.Metadata, error)
	Symbols(context.Context) ([]vm.FunctionDefinition, error)
	Call(ctx context.Context, module vm.Module, function string, inputs []uint64) ([]uint64, error)
}
