package common

import "context"

type Builder interface {
	// Wasm builds a wasm file from the given codeFiles in a temp directory
	Wasm(ctx context.Context, codeFiles ...string) (wasmFile string, err error)
	// Plugin builds a plugin in a temp directory
	Plugin(path string, name string, extraArgs ...string) (string, error)
	// For internal use, contains the tarball of common assets used for building a wasm file
	Fixture() []byte
	// For internal use, returns the name of the Builder
	Name() string
}
