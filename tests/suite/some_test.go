package suite

import (
	"context"
	"fmt"
	"testing"
)

// TODO: Close Server

func TestXxx(t *testing.T) {
	wasmFile, err := Builder().Go().Wasm(context.Background(), "/home/tafkhan/Documents/Work/Taubyte/Repos/vm-orbit/tests/fixtures/build/basic.go")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(wasmFile)
}
