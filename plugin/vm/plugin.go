package vm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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

func (p *vmPlugin) reconnect() error {
	p.proc.Kill()
	return p.connect()
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
				if event.Name == p.filename && (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) {
					if err := p.reload(); err != nil {
						log.Println(err.Error())
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			case <-p.ctx.Done():
				return
			}
		}
	}()
	dir := filepath.Dir(p.filename)

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Handle ma as multi-address
func Load(filename string, ctx context.Context) (vm.Plugin, error) {
	if len(filename) < 1 {
		return nil, errors.New("cannot load plugin from empty filename")
	}

	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("stat `%s` failed with: %w", filename, err)
	}

	p := &vmPlugin{
		filename:  filename,
		instances: make(map[*pluginInstance]interface{}),
	}
	p.ctx, p.ctxC = context.WithCancel(ctx)

	p.connect()
	err := p.watch()
	if err != nil {
		p.ctxC()
		return nil, fmt.Errorf("watch on file `%s` failed with: %w", filename, err)
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
		if err := pI.cleanup(); err != nil {
			return fmt.Errorf("cleanup plugin `%s` failed with: %w", p.name, err)
		}
	}

	if err := p.reconnect(); err != nil {
		return fmt.Errorf("reconnecting plugin `%s` failed with: %w", p.name, err)
	}

	for pI := range p.instances {
		if err := pI.reload(); err != nil {
			return fmt.Errorf("reloading plugin `%s` failed with:%w", p.name, err)
		}
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

	pI := &pluginInstance{
		plugin:   p,
		instance: instance,
	}

	var err error
	if pI.satellite, err = p.getLink(); err != nil {
		return nil, fmt.Errorf("getting link to satelite failed with: %w", err)
	}

	meta, err := pI.satellite.Meta(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("meta failed with: %w", err)
	}

	if len(p.name) < 1 {
		p.name = meta.Name
	}

	return pI, nil
}

func (p *vmPlugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	pI, err := p.new(instance)
	if err != nil {
		return nil, fmt.Errorf("creating new plugin instance for plugin `%s` failed with: %w", p.name, err)
	}

	p.lock.Lock()
	p.instances[pI] = nil
	p.lock.Unlock()

	return pI, nil
}

func (p *vmPlugin) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	var err error
	for pI := range p.instances {
		if _err := pI.close(); _err != nil {
			err = _err
		}
	}

	p.ctxC()
	p.proc.Kill()
	return err
}
