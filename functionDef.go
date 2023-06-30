package orbit

import "github.com/taubyte/go-interfaces/vm"

var _ vm.FunctionDefinition = &functionDefinition{}

type functionDefinition struct {
	name string
	args []vm.ValueType
	rets []vm.ValueType
}

func (f *functionDefinition) Name() string {
	return f.name
}

func (f *functionDefinition) ParamTypes() []vm.ValueType {
	return f.args
}

func (f *functionDefinition) ResultTypes() []vm.ValueType {
	return f.rets
}
