package tests

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/taubyte/vm-orbit/tests/suite"
	"gotest.tools/v3/assert"
)

func TestHelloWorld(t *testing.T) {
	ctx := context.Background()
	testingSuite, err := suite.New(ctx)
	assert.NilError(t, err)

	builder := suite.Builder().Go()

	wd, err := os.Getwd()
	assert.NilError(t, err)

	pluginPath, err := builder.Plugin(path.Join(wd, ".."), "example")
	assert.NilError(t, err)

	err = testingSuite.AttachPluginFromPath(pluginPath)
	assert.NilError(t, err)

	wasmPath, err := builder.Wasm(ctx, path.Join(wd, "fixtures"))
	assert.NilError(t, err)

	module, err := testingSuite.WasmModule(wasmPath)
	assert.NilError(t, err)

	_, err = module.Call(ctx, "helloWorld")
	assert.NilError(t, err)

	module.Debug()
}
