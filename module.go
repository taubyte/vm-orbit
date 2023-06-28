package orbit

import (
	"errors"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

// func (s *GRPCModuleServer) MemoryRead(ctx context.Context, in *proto.ReadRequest) (*proto.ReadReturn, error) {
// 	ret := proto.ReadReturn{}
// 	data, ok := s.module.Memory().Read(in.Offset, in.Size)
// 	if !ok {
// 		ret.Error = proto.IOError_eof
// 	}

// 	ret.Data = data
// 	return &ret, nil
// }

// func (s *GRPCModuleServer) MemoryWrite(ctx context.Context, in *proto.WriteRequest) (*proto.WriteReturn, error) {
// 	ret := proto.WriteReturn{}
// 	if ok := s.module.Memory().Write(in.Offset, in.Data); !ok {
// 		ret.Error = proto.IOError_eof
// 	} else {
// 		ret.Written = uint32(len(in.Data))
// 	}

// 	return &ret, nil
// }

// func (c *GRPCModuleClient) MemoryRead(offset uint32, size uint32) ([]byte, error) {
// 	ret, err := c.client.MemoryRead(c.ctx, &proto.ReadRequest{
// 		Broker: c.brokerId,
// 		Offset: offset,
// 		Size:   size,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if ret.Error != 0 {
// 		return nil, ret.Error.Error()
// 	}

// 	return ret.Data, nil
// }

// func (c *GRPCModuleClient) MemoryWrite(offset uint32, data []byte) (uint32, error) {
// 	ret, err := c.client.MemoryWrite(c.ctx, &proto.WriteRequest{
// 		Broker: c.brokerId,
// 		Offset: offset,
// 		Data:   data,
// 	})
// 	if err != nil {
// 		return 0, err
// 	}

// 	if ret.Error != 0 {
// 		return 0, ret.Error.Error()
// 	}

// 	return ret.Written, nil
// }

func bytesToTypes(valueTypes []vm.ValueType) ([]proto.Type, error) {
	types := make([]proto.Type, len(valueTypes))
	for idx, vt := range valueTypes {
		switch proto.Type(vt) {
		case proto.Type_i32, proto.Type_i64, proto.Type_f32, proto.Type_f64:
			types[idx] = proto.Type(vt)
		default:
			return nil, errors.New("unknown type")
		}
	}

	return types, nil
}

func typesToBytes(valueTypes []proto.Type) ([]vm.ValueType, error) {
	types := make([]vm.ValueType, len(valueTypes))
	for idx, vt := range valueTypes {
		switch vm.ValueType(vt) {
		case vm.ValueTypeF32, vm.ValueTypeF64, vm.ValueTypeI32, vm.ValueTypeI64:
			types[idx] = vm.ValueType(vt)
		default:
			return nil, errors.New("unknown type")
		}
	}

	return types, nil
}
