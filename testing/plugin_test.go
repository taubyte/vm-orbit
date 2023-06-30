package testing

import (
	goCTX "context"
	"fmt"
	"testing"

	"github.com/taubyte/go-interfaces/services/tns/mocks"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/utils/id"
	"github.com/taubyte/vm-orbit/plugin"

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

	// tns.Inject()

	_ctx, err := context.New(
		ctx,
		context.Application(mocksConfig.Application),
		context.Project(mocksConfig.Project),
		context.Resource(mocksConfig.Cid),
		context.Branch(mocksConfig.Branch),
		context.Commit(mocksConfig.Commit),
	)

	if err != nil {
		t.Error(err)
		return
	}

	instance, err := vmService.New(_ctx, vm.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	rt, err := instance.Runtime(nil)
	if err != nil {
		t.Error(err)
		return
	}

	plugin, err := plugin.Load("/home/tafkhan/Documents/Work/Taubyte/Repos/vm-orbit/testing/plugin/plugin")
	if err != nil {
		t.Error(err)
		return
	}

	_, mi, err := rt.Attach(plugin)
	if err != nil {
		t.Error(err)
		return
	}

	mod, err := rt.Module("/file/" + "/home/tafkhan/Documents/Work/Taubyte/Repos/vm-orbit/testing/main.wasm")
	if err != nil {
		t.Error(err)
		return
	}

	for _, fx := range mi.Functions() {
		fmt.Printf("%#v\n", fx)
	}

	fi, err := mod.Function("ping")
	if err != nil {
		t.Error(err)
		return
	}

	ret := fi.Call(ctx)
	fmt.Println(ret.Error())

}
