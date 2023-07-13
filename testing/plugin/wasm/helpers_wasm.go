package lib

import (
	"reflect"
	"strings"

	"github.com/taubyte/go-sdk/utils/codec"
)

// currently using production to build the main.wasm file

//go:wasm-module testing
//export readWritePlus1
func readWritePlus1(*byte, *byte, *byte, uint32, *byte, string, *byte, *byte, uint32, *byte, *uint16, *uint16, *uint32, *uint32, *uint64, *uint64) uint32

//export ping
func ping() {
	var (
		byteVal                byte = 42
		byteValRcv             byte
		bytesSlice             = [][]byte{{42, 43, 44}, {45, 46, 47}, {48, 49, 50}}
		bytesSliceEncoded      []byte
		bytesSliceEncodedSize  uint32
		bytesSliceRcvEncoded   []byte
		bytesSliceRcv          [][]byte
		stringVal              = "hello world"
		stringRcvRaw           = make([]byte, len([]byte(stringVal+"one")))
		stringSlice            = []string{"hello", "world"}
		stringSliceEncoded     []byte
		stringSliceEncodedSize uint32
		stringSliceRcvEncoded  []byte
		stringSliceRcv         []string
		u16                    uint16 = 52
		u16Rcv                 uint16
		u32                    uint32 = 62
		u32Rcv                 uint32
		u64                    uint64 = 72
		u64Rcv                 uint64
	)

	if err := codec.Convert(bytesSlice).To(&bytesSliceEncoded); err != nil {
		panic(err)
	}

	bytesSliceEncodedSize = uint32(len(bytesSliceEncoded))
	bytesSliceRcvEncoded = make([]byte, bytesSliceEncodedSize)

	if err := codec.Convert(stringSlice).To(&stringSliceEncoded); err != nil {
		panic(err)
	}

	stringSliceEncodedSize = uint32(len(stringSliceEncoded))
	stringSliceRcvEncoded = make([]byte, stringSliceEncodedSize)

	readWritePlus1(&byteVal, &byteValRcv, &bytesSliceEncoded[0], bytesSliceEncodedSize, &bytesSliceRcvEncoded[0], stringVal, &stringRcvRaw[0], &stringSliceEncoded[0], stringSliceEncodedSize, &stringSliceRcvEncoded[0], &u16, &u16Rcv, &u32, &u32Rcv, &u64, &u64Rcv)

	if err := codec.Convert(bytesSliceRcvEncoded).To(&bytesSliceRcv); err != nil {
		panic(err)
	}

	if err := codec.Convert(stringSliceRcvEncoded).To(&stringSliceRcv); err != nil {
		panic(err)
	}

	if byteVal != byteVal {
		panic("byte val not the same")
	}

	if !reflect.DeepEqual(bytesSlice, bytesSliceRcv) {
		panic("bytes slice not same")
	}

	if !strings.EqualFold(stringVal, string(stringRcvRaw)) {
		panic("strings not same")
	}

	if !reflect.DeepEqual(stringSlice, stringSliceRcv) {
		panic("string slice not same")
	}

	if u16 != u16Rcv {
		panic("u16 not same")
	}

	if u32 != u32Rcv {
		panic("u32 not same")
	}

	if u64 != u64Rcv {
		panic("u64 not same")
	}
}
