package tests_test

import (
	"os"
	"path"
	"sync"
	"testing"

	"github.com/taubyte/go-interfaces/vm"
	"gotest.tools/v3/assert"

	vmPlugin "github.com/taubyte/vm-orbit/plugin/vm"
)

func init() {
	if err := initializeAssetPaths(); err != nil {
		panic(err)
	}

	if _, err := os.Stat(path.Join(assetDir, "data_helpers.wasm")); err != nil {
		if err = initializeWasm("data_helpers"); err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(assetDir, "basic.wasm")); err != nil {
		if err = initializeWasm("basic"); err != nil {
			panic(err)
		}
	}
}

func TestPlugin(t *testing.T) {
	err := buildPlugin("")
	assert.NilError(t, err)

	instance, ctx, err := newVM()
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wasmFile, plugin := plugin(t, "basic.wasm", ctx)
	defer plugin.Close()

	fi := getFunction(t, wasmFile, rt, plugin)

	checkCall(t, ctx, fi, 5, 5+42)
}

func TestConcurrentPlugin(t *testing.T) {
	err := buildPlugin("")
	assert.NilError(t, err)

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
	wasmFile, plugin := plugin(t, "basic.wasm", ctx)
	defer plugin.Close()

	for i := 0; i < runtimeCount; i++ {
		go func(idx int) {
			defer wg.Done()
			rt := runtimes[idx]
			fi := getFunction(t, wasmFile, rt, plugin)
			checkCall(t, ctx, fi, uint32(idx), uint32(idx)+42)
		}(i)
	}

	wg.Wait()
}

func TestUpdatePlugin(t *testing.T) {
	pluginEvents := vmPlugin.Subscribe(t)
	defer vmPlugin.UnSubscribe(t)

	instance, ctx, err := newVM()
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wasmFile, plugin := plugin(t, "basic.wasm", ctx)
	defer plugin.Close()

	_, _, err = rt.Attach(plugin)
	assert.NilError(t, err)

	err = buildPlugin("update")
	assert.NilError(t, err)

	<-pluginEvents
	mod, err := rt.Module("/file/" + wasmFile)
	assert.NilError(t, err)

	fi, err := mod.Function("ping")
	assert.NilError(t, err)

	checkCall(t, ctx, fi, 5, 5+43)
}

func TestDataHelpers(t *testing.T) {
	err := buildPlugin("")
	assert.NilError(t, err)

	instance, ctx, err := newVM()
	assert.NilError(t, err)
	defer instance.Close()

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)
	defer rt.Close()

	wasmFile, plugin := plugin(t, "data_helpers.wasm", ctx)
	defer plugin.Close()

	fi := getFunction(t, wasmFile, rt, plugin)

	ret := fi.Call(ctx)
	assert.NilError(t, ret.Error())
}

func TestSizeHelpers(t *testing.T) {
	err := buildPlugin("")
	assert.NilError(t, err)

	instance, ctx, err := newVM()
	assert.NilError(t, err)
	defer instance.Close()

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)
	defer rt.Close()

	wasmFile, plugin := plugin(t, "size_helpers.wasm", ctx)
	defer plugin.Close()

	fi := getFunction(t, wasmFile, rt, plugin)

	ret := fi.Call(ctx)
	assert.NilError(t, ret.Error())
}
