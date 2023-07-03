package vm

import (
	"errors"
	"fmt"
	"os/exec"
	"sync"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/link"
)

func newClient(ma string) *plugin.Client {
	return plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: HandShake(),
			Plugins:         link.ClientPluginMap,
			Cmd:             exec.Command(ma),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		},
	)
}

// TODO: Handle ma as multi-address
func Load(ma string) (vm.Plugin, error) {
	if len(ma) < 1 {
		return nil, errors.New("cannot load plugin from empty multi-address")
	}

	return &vmPlugin{
		address: ma,
		client:  newClient(ma),
		lock:    &sync.RWMutex{},
	}, nil
}

func (p *vmPlugin) Name() string {
	return p.name
}

func (p *vmPlugin) Reload() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.client == nil || len(p.address) < 1 {
		return errors.New("cannot reload plugin before load")
	}

	p.client.Kill()
	p.client = newClient(p.address)

	return nil
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

	if err := pluginClient.AttachLock(p.lock); err != nil {
		return nil, fmt.Errorf("attaching mutex lock failed with: %w", err)
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

func (p *vmPlugin) Close() error {
	p.lock.Lock()
	p.client.Kill()
	p.name = ""
	p.lock.Unlock()

	return nil
}
