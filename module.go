package orbit

import (
	"errors"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/proto"
)

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
