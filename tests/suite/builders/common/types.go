package common

import "context"

type Builder interface {
	Wasm(ctx context.Context, codeFiles ...string) (wasmFile string, err error)
	Plugin(path string, name string, extraArgs ...string) (string, error)
	Fixture() []byte
	Name() string
}
