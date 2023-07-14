package goBuilder

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	_ "embed"

	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
	"github.com/taubyte/vm-orbit/tests/suite/builders/fixtures"

	"github.com/otiai10/copy"
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

func (g *goBuilder) Plugin(_path string, name string, extraArgs ...string) (string, error) {
	args := []string{"build", "-o", name}
	args = append(args, extraArgs...)
	cmd := exec.Command("go", args...)
	cmd.Dir = _path
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("building go plugin failed with: %w", err)
	}

	return path.Join(_path, name), nil
}
