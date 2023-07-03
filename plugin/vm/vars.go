package vm

import (
	"github.com/hashicorp/go-plugin"
)

func HandShake() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "VM_ORBIT_SATELLITE",
		MagicCookieValue: "taubyte",
	}
}
