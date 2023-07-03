package link

import (
	"sync"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

type link struct {
	plugin.NetRPCUnsupportedPlugin
}

type GRPCPluginClient struct {
	broker *plugin.GRPCBroker
	client proto.PluginClient

	lock *sync.RWMutex
}

type module struct {
	proto.UnimplementedModuleServer
	module vm.Module
}

var _ vm.FunctionDefinition = &functionDefinition{}

type functionDefinition struct {
	name string
	args []vm.ValueType
	rets []vm.ValueType
}
