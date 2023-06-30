package satellite

import "github.com/hashicorp/go-plugin"

var (
	ServerPluginMap = map[string]plugin.Plugin{
		"satellite": &satellite{},
	}
)
