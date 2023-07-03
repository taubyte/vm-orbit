package vm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-orbit/link"
)

func (p *vmPlugin) connect() (err error) {
	p.proc = plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: HandShake(),
			Plugins:         link.ClientPluginMap,
			Cmd:             exec.Command(p.filename),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		},
	)

	p.client, err = p.proc.Client()
	if err != nil {
		return fmt.Errorf("getting rpc protocol client failed with: %w", err)
	}

	return
}

func (p *vmPlugin) reconnect() {
	p.proc.Kill()

	p.connect()
}

func (p *vmPlugin) watch() error {
	// will panic if any error
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					p.reload()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			case <-p.ctx.Done():
				return
			}
		}
	}()

	err = watcher.Add(p.filename)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Handle ma as multi-address
func Load(filename string) (vm.Plugin, error) {
	if filename == "" {
		return nil, errors.New("cannot load plugin from empty multi-address")
	}

	p := &vmPlugin{
		filename:  filename,
		instances: make(map[*pluginInstance]interface{}),
	}
	p.ctx, p.ctxC = context.WithCancel(context.Background())

	p.connect()
	err := p.watch()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *vmPlugin) Name() string {
	return p.name
}

func (p *vmPlugin) reload() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	for pI := range p.instances {
		pI.cleanup()
	}

	p.reconnect()

	for pI := range p.instances {
		pI.reload()
	}

	return nil
}

func (p *vmPlugin) getLink() (sat Satellite, err error) {
	raw, err := p.client.Dispense("satellite")
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
	p.proc.Kill()
	return nil
}
