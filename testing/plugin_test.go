package testing_test

import (
	goCTX "context"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/taubyte/go-interfaces/services/tns/mocks"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/utils/id"
	"gotest.tools/v3/assert"

	vmPlugin "github.com/taubyte/vm-orbit/plugin/vm"
	fileBE "github.com/taubyte/vm/backend/file"
	"github.com/taubyte/vm/context"
	loader "github.com/taubyte/vm/loaders/wazero"
	resolver "github.com/taubyte/vm/resolvers/taubyte"
	service "github.com/taubyte/vm/service/wazero"
	source "github.com/taubyte/vm/sources/taubyte"
)

func TestPlugin(t *testing.T) {
	instance, ctx, err := newVM()
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wasmFile, plugin := plugin(t, ctx)
	defer plugin.Close()

	checkCall(t, ctx, wasmFile, rt, plugin, 5)
}

func TestConcurrentPlugin(t *testing.T) {
	instance, ctx, err := newVM()
	assert.NilError(t, err)

	runtimeCount := 1
	runtimes := make([]vm.Runtime, runtimeCount)

	for i := 0; i < runtimeCount; i++ {
		runtimes[i], err = instance.Runtime(nil)
		assert.NilError(t, err)
	}

	var wg sync.WaitGroup
	wg.Add(runtimeCount)
	wasmFile, plugin := plugin(t, ctx)
	defer plugin.Close()

	for i := 0; i < runtimeCount; i++ {
		go func(idx int) {
			defer wg.Done()
			rt := runtimes[idx]
			checkCall(t, ctx, wasmFile, rt, plugin, uint32(idx))
		}(i)
	}

	wg.Wait()
}

func plugin(t *testing.T, ctx goCTX.Context) (wasmFile string, plugin vm.Plugin) {
	wd, err := os.Getwd()
	assert.NilError(t, err)

	pluginBinary := path.Join(wd, "plugin", "plugin")
	wasmFile = path.Join(wd, "plugin", "wasm", "main.wasm")
	plugin, err = vmPlugin.Load(pluginBinary, ctx)
	assert.NilError(t, err)

	return
}

func checkCall(t *testing.T, ctx goCTX.Context, wasmFile string, rt vm.Runtime, plugin vm.Plugin, callVal uint32) {
	_, _, err := rt.Attach(plugin)
	assert.NilError(t, err)

	mod, err := rt.Module("/file/" + wasmFile)
	assert.NilError(t, err)

	fi, err := mod.Function("ping")
	assert.NilError(t, err)

	ret := fi.Call(ctx, callVal)
	assert.NilError(t, ret.Error())

	var output uint32
	err = ret.Reflect(&output)
	assert.NilError(t, err)

	assert.Equal(t, output, callVal+42)
}

func newVM() (vm.Instance, goCTX.Context, error) {
	tns := mocks.New()
	rslver := resolver.New(tns)
	ldr := loader.New(rslver, fileBE.New())
	src := source.New(ldr)
	ctx := goCTX.TODO()
	vmService := service.New(ctx, src)

	mocksConfig := mocks.InjectConfig{
		Branch:      "master",
		Commit:      "head_commit",
		Project:     id.Generate(),
		Application: id.Generate(),
		Cid:         id.Generate(),
	}

	_ctx, err := context.New(
		ctx,
		context.Application(mocksConfig.Application),
		context.Project(mocksConfig.Project),
		context.Resource(mocksConfig.Cid),
		context.Branch(mocksConfig.Branch),
		context.Commit(mocksConfig.Commit),
	)
	if err != nil {
		return nil, nil, err
	}

	instance, err := vmService.New(_ctx, vm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return instance, ctx, err
}
