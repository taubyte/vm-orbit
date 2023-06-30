package orbit

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

var _ Satellite = &GRPCPluginClient{}

// this is used by main process to communicate with plugin process
type GRPCPluginClient struct {
	broker *plugin.GRPCBroker
	client proto.PluginClient
}

func (c *GRPCPluginClient) Symbols(ctx context.Context) ([]vm.FunctionDefinition, error) {
	resp, err := c.client.Symbols(ctx, &proto.Empty{})
	if err != nil {
		return nil, err
	}

	funcDefs := make([]vm.FunctionDefinition, len(resp.Functions))
	for idx, function := range resp.Functions {
		args, err := typesToBytes(function.Args)
		if err != nil {
			return nil, fmt.Errorf("getting args failed with: %s", err)
		}

		rets, err := typesToBytes(function.Rets)
		if err != nil {
			return nil, fmt.Errorf("getting returns failed with: %s", err)
		}

		funcDefs[idx] = &functionDefinition{
			name: function.Name,
			args: args,
			rets: rets,
		}
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
		return nil, err
	}
	defer s.Stop()

	if resp.Error != nil {
		if resp.Error.Code != nil {
			return nil, fmt.Errorf("failed with code: %d", *resp.Error.Code)
		}

		return nil, fmt.Errorf("failed with: %s", resp.Error.Message)
	}

	return resp.Rets, nil
}

type module struct {
	proto.UnimplementedModuleServer
	module vm.Module
}

func NewModule(mod vm.Module) proto.ModuleServer {
	return &module{module: mod}
}

func (m *module) MemoryRead(ctx context.Context, req *proto.ReadRequest) (*proto.ReadReturn, error) {
	data, ok := m.module.Memory().Read(req.Offset, req.Size)
	if !ok {
		return nil, errors.New("bitch")
	}

	return &proto.ReadReturn{Data: data}, nil
}

func (m *module) MemoryWrite(ctx context.Context, req *proto.WriteRequest) (*proto.WriteReturn, error) {
	ok := m.module.Memory().Write(req.Offset, req.Data)
	if !ok {
		return nil, errors.New("bitch")
	}

	return &proto.WriteReturn{Written: uint32(len(req.Data))}, nil
}
