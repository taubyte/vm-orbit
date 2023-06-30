package common

import "github.com/taubyte/go-interfaces/vm"

func NewFuncDefinition(name string, args []vm.ValueType, rets []vm.ValueType) vm.FunctionDefinition {
	return &functionDefinition{
		name: name,
		args: args,
		rets: rets,
	}
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
