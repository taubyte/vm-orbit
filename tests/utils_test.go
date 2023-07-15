package tests

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/otiai10/copy"
	"github.com/taubyte/vm-orbit/tests/suite"
	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
)

var (
	wd           string
	fixtureDir   string
	basicWasm    string
	pluginBinary string

	wasmFixtures = []string{"data_helpers", "size_helpers", "basic"}
	builder      common.Builder
	pluginName   = "testPlugin"
)

func initializeAssetPaths() (err error) {
	if wd, err = os.Getwd(); err != nil {
		return
	}

	fixtureDir = path.Join(wd, "fixtures")
	basicWasm = path.Join(fixtureDir, "basic.wasm")
	pluginBinary = path.Join(fixtureDir, pluginName)

	return
}

func initializeBuilder() {
	if builder == nil {
		builder = suite.Builder().Go()
	}
}

func initializePlugin(extraArgs ...string) (err error) {
	initializeBuilder()
	pluginFile, err := builder.Plugin(path.Join(fixtureDir, "plugin"), pluginName, extraArgs...)
	if err != nil {
		return fmt.Errorf("generating plugin failed with: %w", err)
	}

	if err = copy.Copy(pluginFile, path.Join(fixtureDir, pluginName)); err != nil {
		return fmt.Errorf("copying plugin failed with: %w", err)
	}

	return nil
}

func initializeWasm(name string) (err error) {
	wasmFile, err := builder.Wasm(context.TODO(), path.Join(fixtureDir, "_code", name+".go"))
	if err != nil {
		return fmt.Errorf("generating %s.wasm failed with: %w", name, err)
	}

	if err = copy.Copy(wasmFile, path.Join(fixtureDir, name+".wasm")); err != nil {
		return fmt.Errorf("copying %s.wasm failed with: %w", name, err)
	}

	return nil
}
