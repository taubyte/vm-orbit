package satellite

import (
	"reflect"

	"github.com/hashicorp/go-plugin"
)

var (
	ServerPluginMap = map[string]plugin.Plugin{
		"satellite": &satellite{},
	}

	moduleType = reflect.TypeOf((*Module)(nil)).Elem()
)
