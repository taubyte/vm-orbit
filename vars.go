package orbit

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "EXTERNAL_PLUGIN",
	MagicCookieValue: "taubyte",
}

var ClientPluginMap = map[string]plugin.Plugin{
	"satellite": NewLink(),
	//"modulePlugin": &ModulePlugin{},
}

var ServerPluginMap = map[string]plugin.Plugin{
	"satellite": &satellite{},
	//"modulePlugin": &ModulePlugin{},
}
