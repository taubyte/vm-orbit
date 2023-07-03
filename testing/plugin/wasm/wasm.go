package lib

// currently using production to build the main.wasm file

import "fmt"

//go:wasm-module testing
//export add42
func add42(*byte, uint32) uint32

//export ping
func ping(val uint32) uint32 {
	valByteString := []byte(fmt.Sprintf("%d", val))

	sum := add42(&valByteString[0], uint32(len(valByteString)))

	return sum
}
