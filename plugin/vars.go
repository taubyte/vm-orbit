package plugin

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "EXTERNAL_PLUGIN",
	MagicCookieValue: "taubyte",
}
