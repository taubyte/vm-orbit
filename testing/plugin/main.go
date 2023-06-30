package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-plugin"
	orbit "github.com/taubyte/vm-orbit"
)

func hello(ctx context.Context, module orbit.Module, num uint32) uint32 {
	f, _ := os.Create("/tmp/hello.txt")
	defer f.Close()

	fmt.Println(module)
	fmt.Fprintln(f, "The answer is:: ", num)

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
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: orbit.Handshake,
		Plugins: map[string]plugin.Plugin{
			"satellite": orbit.NewSatellite(
				"aladdin",
				exports,
			),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
