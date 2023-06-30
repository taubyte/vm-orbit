package common

import (
	"reflect"

	"github.com/hashicorp/go-plugin"
)

var (
	ModuleType = reflect.TypeOf((*Module)(nil)).Elem()

	Handshake = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "EXTERNAL_PLUGIN",
		MagicCookieValue: "taubyte",
	}
)
