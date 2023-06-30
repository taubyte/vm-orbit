package satellite

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/proto"
)

type satellite struct {
	plugin.NetRPCUnsupportedPlugin

	name    string
	exports map[string]interface{}
}

type GRPCPluginServer struct {
	broker *plugin.GRPCBroker
	proto.UnimplementedPluginServer

	satellite *satellite
}

type moduleLink struct {
	plugin.NetRPCUnsupportedPlugin
	ctx    context.Context
	client proto.ModuleClient
}

type Module interface {
	MemoryRead(offset uint32, size uint32) ([]byte, error)
	MemoryWrite(offset uint32, data []byte) (uint32, error)
}
