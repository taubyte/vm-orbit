package orbit

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

// TODO: Possibly turn this into an interface
type satellite struct {
	plugin.NetRPCUnsupportedPlugin

	name    string
	exports map[string]interface{}
}

func NewSatellite(
	name string,
	exports func() map[string]interface{},
) plugin.Plugin {
	return &satellite{
		name:    name,
		exports: exports(),
	}
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
