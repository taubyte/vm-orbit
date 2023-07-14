package tests

import (
	goCTX "context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	builder "bitbucket.org/taubyte/go-builder"
	"github.com/otiai10/copy"
	"github.com/taubyte/go-interfaces/services/tns/mocks"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/utils/id"
	vmPlugin "github.com/taubyte/vm-orbit/plugin/vm"
	fileBE "github.com/taubyte/vm/backend/file"
	"github.com/taubyte/vm/context"
	loader "github.com/taubyte/vm/loaders/wazero"
	resolver "github.com/taubyte/vm/resolvers/taubyte"
	service "github.com/taubyte/vm/service/wazero"
	source "github.com/taubyte/vm/sources/taubyte"
	"gotest.tools/v3/assert"
)

var (
	wd         string
	assetDir   string
	taubyteDir = ".taubyte"
	goMod      = "go.mod"
	buildDir   string
)

func initializeAssetPaths() (err error) {
	if wd, err = os.Getwd(); err != nil {
		return
	}

	assetDir = path.Join(wd, "assets")
	buildDir = path.Join(assetDir, "build")

	return
}

func goExtension(fileName string) string {
	return fileName + ".go"
}

func initializeWasm(fileName string) error {
	tempDir, err := os.MkdirTemp("/tmp", "*")
	if err != nil {
		return fmt.Errorf("creating temp dir failed with: %w", err)
	}

	goFile := goExtension(fileName)
	if err = copy.Copy(path.Join(buildDir, goFile), path.Join(tempDir, goFile)); err != nil {
		return fmt.Errorf("copying `%s` failed with: %w", goFile, err)
	}

	if err = copy.Copy(path.Join(buildDir, taubyteDir), path.Join(tempDir, taubyteDir)); err != nil {
		return fmt.Errorf("copying taubyteDir failed with: %w", err)
	}

	if err = copy.Copy(path.Join(buildDir, goMod), path.Join(tempDir, goMod)); err != nil {
		return fmt.Errorf("copying go.mod failed with: %w", err)
	}

	_builder, err := builder.New(goCTX.TODO(), tempDir)
	if err != nil {
		return fmt.Errorf("creating new builder failed with: %w", err)
	}

	out, err := _builder.Build()
	if err != nil {
		return fmt.Errorf("builder.Build failed with: %w", err)
	}

	wasmPath := path.Join(assetDir, fileName+".wasm")
	if err := copy.Copy(path.Join(out.OutDir(), "artifact.wasm"), wasmPath); err != nil {
		return fmt.Errorf("copying wasm build failed with: %w", err)
	}

	return nil
}

func plugin(t *testing.T, wasmFileName string, ctx goCTX.Context) (wasmFile string, plugin vm.Plugin) {
	wd, err := os.Getwd()
	assert.NilError(t, err)

	pluginBinary := path.Join(wd, "plugin", "plugin")
	wasmFile = path.Join(wd, "assets", wasmFileName)
	plugin, err = vmPlugin.Load(pluginBinary, ctx)
	assert.NilError(t, err)

	return
}

func getFunction(t *testing.T, wasmFile string, rt vm.Runtime, plugin vm.Plugin) vm.FunctionInstance {
	_, _, err := rt.Attach(plugin)
	assert.NilError(t, err)

	mod, err := rt.Module("/file/" + wasmFile)
	assert.NilError(t, err)

	fi, err := mod.Function("ping")
	assert.NilError(t, err)

	return fi
}

func checkCall(t *testing.T, ctx goCTX.Context, fi vm.FunctionInstance, callVal uint32, expected uint32) {
	ret := fi.Call(ctx, callVal)
	assert.NilError(t, ret.Error())

	var output uint32
	err := ret.Reflect(&output)
	assert.NilError(t, err)

	assert.Equal(t, output, expected)
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

func buildPlugin(buildTag string) error {
	pluginDir := "./plugin"
	args := []string{"build"}
	if len(buildTag) > 1 {
		args = append(args, "-tags", buildTag)
	}

	cmd := exec.Command("go", args...)
	cmd.Dir = pluginDir

	return cmd.Run()
}
