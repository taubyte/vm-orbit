package builders

import (
	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
	goBuilder "github.com/taubyte/vm-orbit/tests/suite/builders/go"
)

// Returns a Go Builder for Wasm File, and Plugin generation
func (b *builders) Go() common.Builder {
	return goBuilder.New()
}
