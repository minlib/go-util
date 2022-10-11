package file

import (
	"os"
	"path/filepath"
)

// Dir get directory in path.
func Dir(path string) string {
	return filepath.Dir(path)
}

// MkdirAll create directory of paths.
func MkdirAll(paths ...string) error {
	for _, path := range paths {
		directory := Dir(path)
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// IsDir returns the path is a directory.
func IsDir(path string) bool {
	if file, err := os.Stat(path); err == nil {
		return file.IsDir()
	}
	return false
}

// Exist returns the file or directory exists.
func Exist(path string) bool {
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		return true
	}
	return false
}
