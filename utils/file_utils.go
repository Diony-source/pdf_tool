package utils

import (
	"fmt"
	"os"
)

// CheckFileExists checks if a file exists
func CheckFileExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	return nil
}
