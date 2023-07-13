package testing_test

import (
	goCTX "context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"testing"

	builder "bitbucket.org/taubyte/go-builder"
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

func TestBuild(t *testing.T) {
	_builder, err := builder.New(goCTX.TODO(), "/home/tafkhan/Documents/Work/Taubyte/test/tb_code_testCustomDomain/functions/ping_pong")
	if err != nil {
		t.Error(err)
		return
	}

	out, err := _builder.Build()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(out.OutDir())
}

func TestHelpers(t *testing.T) {
	err := buildPlugin()
	assert.NilError(t, err)

	instance, ctx, err := newVM()
	assert.NilError(t, err)
	defer instance.Close()

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)
	defer rt.Close()

	wasmFile, plugin := plugin(t, "helpers.wasm", ctx)
	defer plugin.Close()

	fi := getFunction(t, wasmFile, rt, plugin)

	ret := fi.Call(ctx)

	stderr, err := io.ReadAll(rt.Stderr())
	if err == nil {
		fmt.Println("STDERR:", string(stderr))
	}

	stdout, err := io.ReadAll(rt.Stdout())
	if err == nil {
		fmt.Println("STDOUT:", string(stdout))
	}

	fmt.Println("EXTRA OUT:", ret.Error())

}

func TestPlugin(t *testing.T) {
	err := buildPlugin()
	assert.NilError(t, err)

	instance, ctx, err := newVM()
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wasmFile, plugin := plugin(t, "main.wasm", ctx)
	defer plugin.Close()

	fi := getFunction(t, wasmFile, rt, plugin)

	checkCall(t, ctx, fi, 5, 5+42)
}

func TestConcurrentPlugin(t *testing.T) {
	err := buildPlugin()
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
	wasmFile, plugin := plugin(t, "main.wasm", ctx)
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

// TODO: Use build flags instead
func TestUpdatePlugin(t *testing.T) {
	pluginEvents := vmPlugin.Subscribe(t)
	defer vmPlugin.UnSubscribe(t)

	instance, ctx, err := newVM()
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wasmFile, plugin := plugin(t, "main.wasm", ctx)
	defer plugin.Close()

	_, _, err = rt.Attach(plugin)
	assert.NilError(t, err)

	pluginDir := "./plugin"
	mainFile := path.Join(pluginDir, "main.go")
	data, err := os.ReadFile(mainFile)
	assert.NilError(t, err)

	dataString := string(data)

	str := strings.Replace(dataString, "uint32(val) + 42", "uint32(val) + 43", -1)

	err = os.WriteFile(mainFile, []byte(str), 0644)
	defer func() {
		os.WriteFile(mainFile, []byte(dataString), 0644)
		buildPlugin()
	}()

	assert.NilError(t, err)

	err = buildPlugin()
	assert.NilError(t, err)

	// Wait for two checks of waitTillCopy
	<-pluginEvents
	mod, err := rt.Module("/file/" + wasmFile)
	assert.NilError(t, err)

	fi, err := mod.Function("ping")
	assert.NilError(t, err)

	checkCall(t, ctx, fi, 5, 5+43)
}

func plugin(t *testing.T, wasmFileName string, ctx goCTX.Context) (wasmFile string, plugin vm.Plugin) {
	wd, err := os.Getwd()
	assert.NilError(t, err)

	pluginBinary := path.Join(wd, "plugin", "plugin")
	wasmFile = path.Join(wd, "plugin", "wasm", wasmFileName)
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

func buildPlugin() error {
	pluginDir := "./plugin"
	cmd := exec.Command("go", "build")
	cmd.Dir = pluginDir
	return cmd.Run()
}
