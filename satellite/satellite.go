package satellite

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

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
