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
