package suite

import (
	"github.com/taubyte/vm-orbit/tests/suite/builders"
	"github.com/taubyte/vm-orbit/tests/suite/builders/common"
)

type BuildHelper interface {
	Go() common.Builder
}

func Builder() BuildHelper {
	return builders.New()
}
