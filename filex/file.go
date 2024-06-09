package filex

import (
	"os"
	"path/filepath"
)

// Dir get directory in path.
func Dir(path string) string {
	return filepath.Dir(path)
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

// WriteFile write file
func WriteFile(filename string, data string) error {
	if err := MkdirAll(filename); err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(data), os.ModePerm)
}

// ReadFile read file
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// GetSize get file size
func GetSize(path string) (int64, error) {
	if file, err := os.Stat(path); err != nil {
		return 0, err
	} else {
		return file.Size(), nil
	}
}
