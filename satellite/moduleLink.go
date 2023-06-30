package satellite

import (
	"context"

	"github.com/taubyte/vm-orbit/common"
	"github.com/taubyte/vm-orbit/proto"
	"google.golang.org/grpc"
)

func NewModuleLink(ctx context.Context, conn *grpc.ClientConn) common.Module {
	return &moduleLink{ctx: ctx, client: proto.NewModuleClient(conn)}
}

func (p *moduleLink) MemoryRead(offset uint32, size uint32) ([]byte, error) {
	ret, err := p.client.MemoryRead(p.ctx, &proto.ReadRequest{Offset: offset, Size: size})
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (p *moduleLink) MemoryWrite(offset uint32, data []byte) (uint32, error) {
	ret, err := p.client.MemoryWrite(p.ctx, &proto.WriteRequest{Offset: offset, Data: data})
	if err != nil {
		return 0, err
	}
	if ret.Error != proto.IOError_none && ret.Error != proto.IOError_eof {
		return 0, ret.Error.Error()
	}
	return uint32(len(data)), ret.Error.Error()
}
