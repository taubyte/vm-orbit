package testing

import (
	goCTX "context"
	"fmt"
	"io"
	"os"
	"path"
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
	assert.NilError(t, err)

	instance, err := vmService.New(_ctx, vm.Config{})
	assert.NilError(t, err)

	rt, err := instance.Runtime(nil)
	assert.NilError(t, err)

	wd, err := os.Getwd()
	assert.NilError(t, err)

	pluginBinary := path.Join(wd, "plugin", "plugin")
	plugin, err := vmPlugin.Load(pluginBinary, ctx)
	assert.NilError(t, err)

	_, _, err = rt.Attach(plugin)
	assert.NilError(t, err)

	wasmFile := path.Join(wd, "main.wasm")
	mod, err := rt.Module("/file/" + wasmFile)
	assert.NilError(t, err)

	fi, err := mod.Function("ping")
	assert.NilError(t, err)

	ret := fi.Call(ctx)
	defer os.Remove("/tmp/hello.txt")
	assert.NilError(t, ret.Error())
	fmt.Println(ret.Error())

	file, err := os.Open("/tmp/hello.txt")
	defer file.Close()
	assert.NilError(t, err)

	data, err := io.ReadAll(file)
	assert.NilError(t, err)
	fmt.Println(string(data))

	assert.DeepEqual(t, string(data), "The answer is: 2\n")

}
