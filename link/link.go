package link

import (
	"context"
	"errors"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

func (p *link) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("can't create a satellite (link server) from main process")
}

func (p *link) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCPluginClient{
		client: proto.NewPluginClient(c),
		broker: broker,
	}, nil
}
