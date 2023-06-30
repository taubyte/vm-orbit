package plugin

import "github.com/hashicorp/go-plugin"

var handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "EXTERNAL_PLUGIN",
	MagicCookieValue: "taubyte",
}
