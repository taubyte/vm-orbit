package plugin

import (
	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/common"
)

type pluginInstance struct {
	plugin   *vmPlugin
	instance vm.Instance
	iface    common.Satellite
}

type vmPlugin struct {
	client  *plugin.Client
	address string
	name    string
}
