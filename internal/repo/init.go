package repo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const defaultHEAD = "ref: refs/heads/main\n"

func Init() (bool, error) {
	return InitAt(".")
}

func InitAt(root string) (bool, error) {
	join := func(elem ...string) string {
		return filepath.Join(append([]string{root}, elem...)...)
	}

	vexDir := join(VexDir)

	// Assuming the current dir is Fresh
	freshDir := true

	// Now, check if that dir exist or not
	if fil, err := os.Stat(vexDir); err == nil {
		if !fil.IsDir() {
			return false, fmt.Errorf("%s exists and is not a directory", vexDir)
		}

		freshDir = false
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("stat %s: %w", vexDir, err)
	}

	// Directories to made
	targetDirs := []string{
		vexDir,
		join(VexDir, "objects"),
		join(VexDir, "refs", "heads"),
		join(VexDir, "refs", "tags"),
		join(VexDir, "logs", "refs", "heads"),
		join(VexDir, "info"),
		join(VexDir, "hooks"),
	}

	for _, dir := range targetDirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return false, fmt.Errorf("create %s: %w", dir, err)
		}
	}

	return freshDir, nil
}
