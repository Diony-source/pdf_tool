package utils

import (
	"os"
	"path/filepath"
)

// FindFileInSystem searches for a file in the system starting from the given root directory
func FindFileInSystem(rootDir, targetFile string) (string, error) {
	var result string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Name() == targetFile {
			result = path
			return filepath.SkipDir // Stop further walking
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", os.ErrNotExist
	}
	return result, nil
}
