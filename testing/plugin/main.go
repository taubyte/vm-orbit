package main

import (
	"context"
	"strconv"

	"github.com/taubyte/vm-orbit/plugin"
	"github.com/taubyte/vm-orbit/satellite"
)

type tester struct{}

func (t *tester) W_add42(ctx context.Context, module satellite.Module, stringPtr uint32, lenPtr uint32) uint32 {
	data, err := module.MemoryRead(stringPtr, lenPtr)
	if err != nil {
		panic(err)
	}

	val, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}

	return uint32(val) + 42
}

func main() {
	plugin.Export("testing", &tester{})
}
