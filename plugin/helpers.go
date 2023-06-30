package plugin

import (
	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/satellite"
)

func Serve(moduleName string, exports func() map[string]interface{}) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		Plugins: map[string]plugin.Plugin{
			"satellite": satellite.New(moduleName, exports),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
