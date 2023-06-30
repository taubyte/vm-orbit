package orbit

import (
	"reflect"

	"github.com/hashicorp/go-plugin"
)

var (
	// TODO: Move Cookie K/V to specs
	Handshake = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "EXTERNAL_PLUGIN",
		MagicCookieValue: "taubyte",
	}

	ClientPluginMap = map[string]plugin.Plugin{
		"satellite": &link{},
	}

	ServerPluginMap = map[string]plugin.Plugin{
		"satellite": &satellite{},
	}

	moduleType = reflect.TypeOf((*Module)(nil)).Elem()
)
