package plugin

import (
	"reflect"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/taubyte/vm-orbit/plugin/vm"
	"github.com/taubyte/vm-orbit/satellite"
)

func Serve(moduleName string, structure interface{}) {
	ServeRaw(moduleName, exports(structure))
}

func ServeRaw(moduleName string, exports func() map[string]interface{}) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: vm.HandShake(),
		Plugins: map[string]plugin.Plugin{
			"satellite": satellite.New(moduleName, exports),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

func exports(structure interface{}) func() map[string]interface{} {
	m := reflect.ValueOf(structure)
	mT := reflect.TypeOf(structure)

	return func() map[string]interface{} {
		exports := make(map[string]interface{}, 0)
		for i := 0; i < m.NumMethod(); i++ {
			mt := m.Method(i)
			mtT := mT.Method(i)
			if strings.HasPrefix(mtT.Name, "W_") {
				exports[mtT.Name[2:]] = mt.Interface()
			}
		}

		return exports
	}
}
