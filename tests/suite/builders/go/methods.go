package goBuilder

import (
	"context"
	"path"

	_ "embed"

	"github.com/otiai10/copy"
	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
	"github.com/taubyte/vm-orbit/tests/suite/builders/fixtures"
)

type goBuilder struct{}

func New() *goBuilder {
	return &goBuilder{}
}

func (g *goBuilder) Wasm(ctx context.Context, codeFiles ...string) (wasmFile string, err error) {
	tempDir, err := common.Fixtures(g)
	if err != nil {
		return
	}

	// TODO: This may become generic
	for _, codeFile := range codeFiles {
		if err = copy.Copy(codeFile, path.Join(tempDir, path.Base(codeFile))); err != nil {
			return
		}
	}

	return common.Wasm(ctx, g, tempDir)
}

func (g *goBuilder) Fixture() []byte {
	return fixtures.GoFixture
}

func (g *goBuilder) Name() string {
	return "go"
}

func (g *goBuilder) Plugin() {}
