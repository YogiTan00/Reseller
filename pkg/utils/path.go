package utils

import (
	"os"
	"path/filepath"
)

func Path(fileName string) (string, error) {
	// Get the directory of the executable
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exePath)

	// Construct the full path to the template file
	fullPath := filepath.Join(dir, fileName)
	return fullPath, nil
}
