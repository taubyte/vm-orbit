package vm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/multiformats/go-multihash"
)

func prepFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("opening file %s failed with: %w", path, err)
	}
	defer f.Close()

	name := filepath.Base(path)

	file, err := os.CreateTemp("/tmp", name)
	if err != nil {
		//verbose
		return "", err
	}

	defer file.Close()
	if _, err := io.Copy(file, f); err != nil {
		//verbose
		return "", err
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		// verbose
		return "", err
	}

	mh, err := multihash.SumStream(f, multihash.SHA2_256, -1)
	if err != nil {
		return "", fmt.Errorf("multihashing %s failed with: %w", f.Name(), err)
	}

	err = os.WriteFile(fmt.Sprintf("%s.sha256", path), []byte(mh.B58String()), 0755)
	if err != nil {
		return "", fmt.Errorf("writing to %s.sha256 failed with: %w", path, err)
	}

	return file.Name(), nil
}
