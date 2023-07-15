package suite

import (
	"context"
	"fmt"

	"github.com/taubyte/go-interfaces/services/tns/mocks"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/utils/id"
	vmPlugin "github.com/taubyte/vm-orbit/plugin/vm"
	"github.com/taubyte/vm-orbit/tests/suite/builders"
	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
	fileBE "github.com/taubyte/vm/backend/file"
	vmContext "github.com/taubyte/vm/context"
	loader "github.com/taubyte/vm/loaders/wazero"
	resolver "github.com/taubyte/vm/resolvers/taubyte"
	service "github.com/taubyte/vm/service/wazero"
	source "github.com/taubyte/vm/sources/taubyte"
)

type buildHelper interface {
	Go() common.Builder
}

func Builder() buildHelper {
	return builders.New()
}

type suite struct {
	ctx      context.Context
	ctxC     context.CancelFunc
	instance vm.Instance
	runtime  vm.Runtime
}

type module struct {
	suite *suite
	mI    vm.ModuleInstance
}

func New(ctx context.Context) (*suite, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var ctxC context.CancelFunc
	ctx, ctxC = context.WithCancel(ctx)

	tns := mocks.New()
	rslver := resolver.New(tns)
	ldr := loader.New(rslver, fileBE.New())
	src := source.New(ldr)
	vmService := service.New(ctx, src)

	vmCtx, err := vmContext.New(
		ctx,
		vmContext.Application(id.Generate()),
		vmContext.Project(id.Generate()),
		vmContext.Resource(id.Generate()),
		vmContext.Branch("master"),
		vmContext.Commit("head_commit"),
	)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("creating new vm context failed with: %w", err)
	}

	instance, err := vmService.New(vmCtx, vm.Config{})
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("creating new vm instance failed with: %w", err)
	}

	rt, err := instance.Runtime(nil)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("creating new vm runtime failed with: %w", err)
	}

	return &suite{
		instance: instance,
		runtime:  rt,
		ctx:      ctx,
		ctxC:     ctxC,
	}, nil
}

func (s *suite) AttachPlugin(plugin vm.Plugin) error {
	if _, _, err := s.runtime.Attach(plugin); err != nil {
		return fmt.Errorf("attaching plugin `%s` failed with: %w", plugin.Name(), err)
	}

	return nil
}

func (s *suite) AttachPluginFromPath(filename string) error {
	plugin, err := vmPlugin.Load(filename, s.ctx)
	if err != nil {
		return fmt.Errorf("loading plugin `%s` failed with: %w", filename, err)
	}

	if _, _, err = s.runtime.Attach(plugin); err != nil {
		return fmt.Errorf("attaching plugin `%s` failed with: %w", plugin.Name(), err)
	}

	return nil
}

func (s *suite) Close() {
	s.runtime.Close()
	s.instance.Close()
	s.ctxC()
}

func (s *suite) WasmModule(filename string) (*module, error) {
	mod, err := s.runtime.Module("/file/" + filename)
	if err != nil {
		return nil, fmt.Errorf("creating new module instance failed with: %w", err)
	}

	return &module{
		suite: s,
		mI:    mod,
	}, nil
}

func (m *module) Call(ctx context.Context, function string, args ...interface{}) (vm.Return, error) {
	fI, err := m.mI.Function(function)
	if err != nil {
		return nil, fmt.Errorf("getting function `%s` failed with: ", err)
	}

	ret := fI.Call(ctx, args...)
	if ret.Error() != nil {
		return nil, fmt.Errorf("calling `%s` failed with: %w", function, ret.Error())
	}

	return ret, nil
}

func GoBuildTags(tags ...string) []string {
	args := []string{"-tags"}
	args = append(args, tags...)
	return args
}
