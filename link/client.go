package link

import (
	"context"
	"fmt"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/common"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

func clientErr(msg string, args ...any) error {
	return fmt.Errorf("[]client "+msg, args...)
}

func (c *GRPCPluginClient) Symbols(ctx context.Context) ([]vm.FunctionDefinition, error) {
	resp, err := c.client.Symbols(ctx, &proto.Empty{})
	if err != nil {
		return nil, clientErr("calling symbols failed with: %w", err)
	}

	funcDefs := make([]vm.FunctionDefinition, len(resp.Functions))
	for idx, function := range resp.Functions {
		args, err := common.TypesToBytes(function.Args)
		if err != nil {
			return nil, clientErr("getting arg types failed with: %s", err)
		}

		rets, err := common.TypesToBytes(function.Rets)
		if err != nil {
			return nil, clientErr("getting return types failed with: %s", err)
		}

		funcDefs[idx] = common.NewFuncDefinition(function.Name, args, rets)
	}

	return funcDefs, nil
}

func (c *GRPCPluginClient) Meta(ctx context.Context) (*proto.Metadata, error) {
	return c.client.Meta(ctx, &proto.Empty{})
}

func (c *GRPCPluginClient) Call(ctx context.Context, module vm.Module, function string, inputs []uint64) ([]uint64, error) {
	moduleServer := NewModule(module)

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)

		proto.RegisterModuleServer(s, moduleServer)

		return s
	}

	brokerID := c.broker.NextId()
	go c.broker.AcceptAndServe(brokerID, serverFunc)

	resp, err := c.client.Call(ctx, &proto.CallRequest{
		Broker:   brokerID,
		Function: function,
		Inputs:   inputs,
	})
	if err != nil {
		return nil, clientErr("calling `%s/%s` failed with: %w", module.Name(), function, err)
	}
	defer s.Stop()

	if resp.Error != nil {
		if resp.Error.Code != nil {
			return nil, clientErr("`%s/%s` failed with code: %d", module.Name(), function, *resp.Error.Code)
		}

		return nil, fmt.Errorf("`%s/%s`failed with: %s", module.Name(), function, resp.Error.Message)
	}

	return resp.Rets, nil
}
