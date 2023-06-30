package link

import "github.com/hashicorp/go-plugin"

var (
	ClientPluginMap = map[string]plugin.Plugin{
		"satellite": &link{},
	}
)
