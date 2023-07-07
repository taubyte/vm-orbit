package vm

import (
	"fmt"
	"os"

	"github.com/multiformats/go-multihash"
)

func hashFileContent(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("opening file %s failed with: %w", filepath, err)
	}
	defer f.Close()

	mh, err := multihash.SumStream(f, multihash.SHA2_256, -1)
	if err != nil {
		return fmt.Errorf("multihashing %s failed with: %w", f.Name(), err)
	}

	err = os.WriteFile(fmt.Sprintf("%s.sha256", filepath), []byte(mh.B58String()), 0755)
	if err != nil {
		return fmt.Errorf("writing to %s.sha256 failed with: %w", filepath, err)
	}

	return nil
}
