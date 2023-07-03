package main

import (
	"context"
	"fmt"
	"os"

	"github.com/taubyte/vm-orbit/plugin"
	"github.com/taubyte/vm-orbit/satellite"
)

type tester struct{}

func (t *tester) W_hello(ctx context.Context, module satellite.Module, num uint32) uint32 {
	f, _ := os.Create("/tmp/hello.txt")
	defer f.Close()

	fmt.Println(module)
	fmt.Fprintln(f, "The answer is:", num)

	return 0
}

func (t *tester) W_sum(a, b int64) int64 {
	return a + b
}

func main() {
	plugin.Export("aladdin", &tester{})
}
