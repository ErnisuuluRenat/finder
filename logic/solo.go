package logic

import (
	"fmt"
	"os"
	"path/filepath"
)

func Solo(fileName string) (string, error) {
	var foundFilePath string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == fileName {
			foundFilePath = path
			return filepath.SkipDir // Stop walking through the directory
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if foundFilePath == "" {
		return "", fmt.Errorf("file not found: %s", fileName)
	}
	return foundFilePath, nil
}
