package plugin

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/link"
)

// TODO: Handle ma as multi-address
func Load(ma string) (vm.Plugin, error) {
	p := &vmPlugin{address: ma}
	p.client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshake,
		Plugins:         link.ClientPluginMap,
		Cmd:             exec.Command(ma),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})

	return p, nil
}

func (p *vmPlugin) Name() string {
	return p.name
}

func (p *vmPlugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	rpcClient, err := p.client.Client()
	if err != nil {
		return nil, fmt.Errorf("getting rpc protocol client failed with: %w", err)
	}

	go func() {
		<-instance.Context().Context().Done()
		rpcClient.Close()
	}()

	raw, err := rpcClient.Dispense("satellite")
	if err != nil {
		return nil, fmt.Errorf("getting satellite failed with: %w", err)
	}

	pluginClient, ok := raw.(Satellite)
	if !ok {
		return nil, errors.New("satellite is not a plugin")
	}

	meta, err := pluginClient.Meta(instance.Context().Context())
	if err != nil {
		return nil, fmt.Errorf("meta failed with: %w", err)
	}

	p.name = meta.Name

	pI := &pluginInstance{
		satellite: pluginClient,
		plugin:    p,
		instance:  instance,
	}

	return pI, nil
}
