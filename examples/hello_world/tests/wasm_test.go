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

	// create a testing suite so we can quickly test our plugin
	testingSuite, err := suite.New(ctx)
	assert.NilError(t, err)

	// create a goBuilder used to build plugins and wasm
	builder := suite.Builder().Go()

	wd, err := os.Getwd()
	assert.NilError(t, err)

	// build the plugin from the parent directory with our main.go with the plugin export
	pluginPath, err := builder.Plugin(path.Join(wd, ".."), "example")
	assert.NilError(t, err)

	// Attaches plugin to our testing suite from the path resolved by builder.Plugin()
	err = testingSuite.AttachPluginFromPath(pluginPath)
	assert.NilError(t, err)

	// build a wasm file from our fixture go file
	wasmPath, err := builder.Wasm(ctx, path.Join(wd, "_fixtures", "dfunc.go"))
	assert.NilError(t, err)

	// get get the wasm module from our wasm file
	module, err := testingSuite.WasmModule(wasmPath)
	assert.NilError(t, err)

	// call our function "helloWorld" from our wasm module
	_, err = module.Call(ctx, "helloWorld")
	assert.NilError(t, err)

	// Prints stdOut and stdErr from our runtime
	module.Debug()
}