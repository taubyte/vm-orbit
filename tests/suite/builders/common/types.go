package common

import "context"

type Builder interface {
	Wasm(ctx context.Context, codeFiles ...string) (wasmFile string, err error)
	Plugin()
	Fixture() []byte
	Name() string
}
