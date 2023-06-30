package main

import (
	"context"
	"fmt"
	"os"

	"github.com/taubyte/vm-orbit/plugin"
	"github.com/taubyte/vm-orbit/satellite"
)

func hello(ctx context.Context, module satellite.Module, num uint32) uint32 {
	f, _ := os.Create("/tmp/hello.txt")
	defer f.Close()

	fmt.Println(module)
	fmt.Fprintln(f, "The new answer is: ", num)

	return 0
}

func sum(a, b int64) int64 {
	return a + b
}

func exports() map[string]interface{} {
	return map[string]interface{}{
		"hello": hello,
		"sum":   sum,
	}
}

func main() {
	plugin.Serve("aladdin", exports)
}
