package plugin

import (
	"fmt"
	"reflect"
	"strings"
)

func Exports(structure interface{}) (func() map[string]interface{}, error) {
	m := reflect.ValueOf(structure)
	mT := reflect.TypeOf(structure)

	if m.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("expected pointer")
	}

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
	}, nil
}
