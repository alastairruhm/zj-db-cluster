package utils

import (
	"os"
	"path/filepath"
)

// Check if a file or directory exists.
func Exist(_path string) (bool, error) {
	_, err := os.Stat(_path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WriteFile ...
func WriteFile(_path string, file string, text string) error {
	f, err := os.Create(filepath.Join(_path, file))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}
