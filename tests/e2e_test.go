package tests

import (
	"context"
	"os"
	"path"
	"sync"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/taubyte/go-interfaces/vm"
	vmPlugin "github.com/taubyte/vm-orbit/plugin/vm"
	"github.com/taubyte/vm-orbit/tests/suite"
)

func init() {
	if err := initializeAssetPaths(); err != nil {
		panic(err)
	}

	if _, err := os.Stat(path.Join(fixtureDir, pluginName)); err != nil {
		if err = initializePlugin(); err != nil {
			panic(err)
		}
	}

	for _, fixture := range wasmFixtures {
		if _, err := os.Stat(path.Join(fixtureDir, fixture+".wasm")); os.IsNotExist(err) {
			if err = initializeWasm(fixture); err != nil {
				panic(err)
			}
		}
	}
}

func basicCall(t *testing.T, plugin vm.Plugin, wasmModule string, args ...interface{}) vm.Return {
	testingSuite, err := suite.New(context.Background())
	assert.NilError(t, err)

	err = testingSuite.AttachPlugin(plugin)
	assert.NilError(t, err)

	module, err := testingSuite.WasmModule(wasmModule)
	assert.NilError(t, err)

	ret, err := module.Call(context.TODO(), "ping", args...)
	assert.NilError(t, err)

	return ret
}

func testReturn(t *testing.T, ret vm.Return, expected uint32) {
	var retVal uint32
	err := ret.Reflect(&retVal)
	assert.NilError(t, err)

	assert.Equal(t, retVal, expected)
}

func TestPlugin(t *testing.T) {
	plugin, err := vmPlugin.Load(pluginBinary, context.Background())
	assert.NilError(t, err)
	ret := basicCall(t, plugin, basicWasm, 5)
	testReturn(t, ret, 47)
}

func TestConcurrentPlugin(t *testing.T) {
	runtimeCount := 5
	plugin, err := vmPlugin.Load(pluginBinary, context.Background())
	assert.NilError(t, err)

	var wg sync.WaitGroup
	wg.Add(runtimeCount)
	for i := 0; i < runtimeCount; i++ {
		go func(val uint32) {
			ret := basicCall(t, plugin, basicWasm, 5)
			testReturn(t, ret, 47)
			wg.Done()
		}(uint32(i))
	}

	wg.Wait()
}

func TestUpdatePlugin(t *testing.T) {
	pluginEvents := vmPlugin.Subscribe(t)
	defer vmPlugin.UnSubscribe(t)

	testingSuite, err := suite.New(context.Background())
	assert.NilError(t, err)

	err = testingSuite.AttachPluginFromPath(pluginBinary)
	assert.NilError(t, err)

	err = initializePlugin(suite.GoBuildTags("update")...)
	assert.NilError(t, err)
	defer initializePlugin()

	<-pluginEvents
	module, err := testingSuite.WasmModule(basicWasm)
	assert.NilError(t, err)

	callVal := uint32(5)
	ret, err := module.Call(context.Background(), "ping", callVal)
	assert.NilError(t, err)

	testReturn(t, ret, 48)
}

func TestDataHelpers(t *testing.T) {
	plugin, err := vmPlugin.Load(pluginBinary, context.Background())
	assert.NilError(t, err)

	basicCall(t, plugin, path.Join(fixtureDir, "data_helpers.wasm"))
}

func TestSizeHelpers(t *testing.T) {
	plugin, err := vmPlugin.Load(pluginBinary, context.Background())
	assert.NilError(t, err)

	basicCall(t, plugin, path.Join(fixtureDir, "size_helpers.wasm"))
}
