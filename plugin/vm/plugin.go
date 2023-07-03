package vm

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/link"
)

func connect(p *vmPlugin) *vmPlugin {
	p.client = plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: HandShake(),
			Plugins:         link.ClientPluginMap,
			Cmd:             exec.Command(p.address),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		},
	)
	return p
}

func reconnect(p *vmPlugin) *vmPlugin {
	p.client.Kill()
	return connect(p)
}

// TODO: Handle ma as multi-address
func Load(ma string) (vm.Plugin, error) {
	if len(ma) < 1 {
		return nil, errors.New("cannot load plugin from empty multi-address")
	}

	p := &vmPlugin{
		address:   ma,
		instances: make(map[*pluginInstance]interface{}),
	}
	p.ctx, p.ctxC = context.WithCancel(context.Background())

	return connect(p), nil
}

func (p *vmPlugin) Name() string {
	return p.name
}

func (p *vmPlugin) Reload() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	reconnect(p)

	for pI := range p.instances {
		pI.reload()
	}

	return nil
}

func (p *vmPlugin) getLink() (sat Satellite, err error) {
	rpcClient, err := p.client.Client()
	if err != nil {
		return nil, fmt.Errorf("getting rpc protocol client failed with: %w", err)
	}

	raw, err := rpcClient.Dispense("satellite")
	if err != nil {
		return nil, fmt.Errorf("getting satellite failed with: %w", err)
	}

	if sat, _ = raw.(Satellite); sat == nil {
		return nil, errors.New("satellite is not a plugin")
	}

	return sat, nil
}

func (p *vmPlugin) new(instance vm.Instance) (*pluginInstance, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	var err error
	pI := &pluginInstance{
		plugin:   p,
		instance: instance,
	}

	if pI.satellite, err = p.getLink(); err != nil {
		return nil, fmt.Errorf("getting link to satelite failed with: %w", err)
	}

	meta, err := pI.satellite.Meta(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("meta failed with: %w", err)
	}

	if p.name == "" {
		p.name = meta.Name
	}

	return pI, nil
}

func (p *vmPlugin) New(instance vm.Instance) (vm.PluginInstance, error) {

	pI, err := p.new(instance)
	if err != nil {
		return nil, err
	}

	p.lock.Lock()
	p.instances[pI] = nil
	p.lock.Unlock()

	return pI, nil
}

func (p *vmPlugin) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	for pI := range p.instances {
		pI.close()
	}

	p.ctxC()
	p.client.Kill()
	return nil
}
