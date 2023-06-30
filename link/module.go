package link

import (
	"context"
	"errors"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

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
