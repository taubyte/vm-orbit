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

	file, err := os.CreateTemp("/tmp", name+"-")
	if err != nil {
		return "", fmt.Errorf("creating temp file %s failed with: %w", name, err)
	}

	if err = file.Chmod(0755); err != nil {
		return "", fmt.Errorf("chmod 0755 on %s failed with: %w", file.Name(), err)
	}

	defer file.Close()
	if _, err := io.Copy(file, f); err != nil {
		return "", fmt.Errorf("copying from %s to %s failed with: %w", f.Name(), file.Name(), err)
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("seeking start in file %s failed with: %w", f.Name(), err)
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
