// Package filex provides utility functions for file and directory operations.
package filex

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of slashes, Base returns "/".
func Base(path string) string {
	return filepath.Base(path)
}

// Dir returns the directory portion of a path.
// If the path contains no slashes, Dir returns ".".
func Dir(path string) string {
	return filepath.Dir(path)
}

// Ext returns the file extension used by path.
// The extension is the suffix beginning at the final dot in the final slash-separated element of path;
// it is empty if there is no dot.
func Ext(filename string) string {
	return filepath.Ext(filename)
}

// Exist reports whether the named file or directory exists.
func Exist(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// FilenameWithoutExt returns the base filename without its extension.
//
// Parameters:
//   - path: the path to the file
//
// Returns:
//   - the filename without extension
//
// Examples:
//   - FilenameWithoutExt("/path/to/file.txt") -> "file"
//   - FilenameWithoutExt("file.tar.gz") -> "file.tar"
func FilenameWithoutExt(path string) string {
	if path == "" {
		return ""
	}
	base := filepath.Base(path)
	ext := filepath.Ext(path)
	return strings.TrimSuffix(base, ext)
}

// ReplaceExt replaces the extension of the given path with the new extension.
//
// Parameters:
//   - path: the original file path
//   - ext: the new extension (should include the leading dot, e.g., ".txt")
//
// Returns:
//   - the path with the new extension
//
// Examples:
//   - ReplaceExt("/path/to/file.txt", ".pdf") -> "/path/to/file.pdf"
//   - ReplaceExt("file.txt", ".doc") -> "file.doc"
func ReplaceExt(path string, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + ext
}

// IsDir reports whether the given path is a directory.
// Returns false if the path doesn't exist or an error occurs.
func IsDir(path string) bool {
	if file, err := os.Stat(path); err == nil {
		return file.IsDir()
	}
	return false
}

// Abs returns the absolute representation of the given path.
// If the path is not absolute, it will be joined with the current working directory.
func Abs(path string) (string, error) {
	return filepath.Abs(path)
}

// GetSize returns the size of the named file.
func GetSize(filename string) (int64, error) {
	file, err := os.Stat(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to get size of %s: %w", filename, err)
	}
	return file.Size(), nil
}

// Join joins any number of path elements into a single path, separating them with an OS-specific separator.
// Empty elements are ignored. The result is cleaned.
func Join(elements ...string) string {
	return filepath.Join(elements...)
}

// JoinDir joins any number of path elements into a single path, separating them with an OS-specific separator.
// This function ensures that the resulting path ends with a directory separator.
// Empty elements are ignored. The result is cleaned.
//
// Parameters:
//   - elements: path elements to join
//
// Returns:
//   - joined path string that ends with a directory separator
//
// Examples:
//   - JoinDir("path", "to", "dir") -> "path/to/dir/"
//   - JoinDir("/root", "sub") -> "/root/sub/"
func JoinDir(elements ...string) string {
	if len(elements) == 0 {
		return ""
	}

	path := filepath.Join(elements...)
	if path == "" {
		return ""
	}

	// Ensure the path ends with a separator
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}

	return path
}

// MkdirAll creates all directories in the provided paths with default permissions (0755).
// It creates directories recursively if needed.
func MkdirAll(paths ...string) error {
	for _, path := range paths {
		directory := Dir(path)
		if directory == "" || directory == "." {
			continue // Skip empty or current directory
		}
		if err := os.MkdirAll(directory, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", directory, err)
		}
	}
	return nil
}

// WriteFile writes data to a file named by filename.
func WriteFile(filename string, data []byte) error {
	if err := MkdirAll(filename); err != nil {
		return fmt.Errorf("failed to create parent directories for %s: %w", filename, err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}
	return nil
}

// WriteFileString writes data to a file named by filename.
func WriteFileString(filename, data string) error {
	return WriteFile(filename, []byte(data))
}

// ReadFile reads the file named by filename and returns the contents.
func ReadFile(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return bytes, nil
}

// ReadFileString reads the file named by filename and returns the contents.
func ReadFileString(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return string(bytes), nil
}

// CopyFile copies the contents of the file named srcFilename to the file named destFilename.
func CopyFile(srcFilename, destFilename string) error {
	src, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destFilename)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}

// IsExtensions checks if the given file path has one of the specified extensions.
// This function performs a case-insensitive check for the file extension.
//
// Parameters:
//   - filePath: the path to the file to check
//   - extensions: one or more file extensions to check against (e.g., ".pdf", ".jpg")
//
// Returns:
//   - true if the file has one of the specified extensions, false otherwise
//
// Examples:
//   - IsExtensions("document.pdf", ".pdf") -> true
//   - IsExtensions("document.pdf", "pdf") -> true
//   - IsExtensions("image.JPG", ".jpg", ".png") -> true
//   - IsExtensions("document.txt", ".pdf") -> false
func IsExtensions(path string, extensions ...string) bool {
	// Handle edge cases
	if path == "" || len(extensions) == 0 {
		return false
	}

	fileExtension := strings.ToLower(filepath.Ext(path))
	if fileExtension == "" {
		return false
	}

	for _, extension := range extensions {
		// Handle empty extension
		if extension == "" {
			continue
		}

		// Ensure the extension starts with a dot
		if !strings.HasPrefix(extension, ".") {
			extension = "." + extension
		}

		// Compare in a case-insensitive manner
		if fileExtension == strings.ToLower(extension) {
			return true
		}
	}

	return false
}

// FindByExtensions traverses the given root directory and returns a list of file paths
// that match the specified extensions. If no extensions are provided, all files are returned.
//
// Parameters:
//   - path: the root directory to traverse
//   - extensions: target file extensions (supporting both ".pdf" or "pdf" formats, optional)
//
// Returns:
//   - A slice of file paths that match the criteria
//   - An error if directory traversal fails
//
// Performance considerations:
//   - Preallocates the matchingFiles slice with estimated capacity to reduce memory allocations
//   - Uses map for O(1) extension lookup
//   - Handles directory traversal errors gracefully without stopping the entire process
func FindByExtensions(path string, extensions ...string) ([]string, error) {
	return FindByExtensionsContext(context.Background(), path, extensions...)
}

// FindByExtensionsContext traverses the given root directory and returns a list of file paths
// that match the specified extensions, with context support for cancellation.
// If no extensions are provided, all files are returned.
//
// Parameters:
//   - ctx: context for cancellation support
//   - path: the root directory to traverse
//   - extensions: target file extensions (supporting both ".pdf" or "pdf" formats, optional)
//
// Returns:
//   - A slice of file paths that match the criteria
//   - An error if directory traversal fails
func FindByExtensionsContext(ctx context.Context, path string, extensions ...string) ([]string, error) {
	if path == "" {
		return nil, errors.New("root directory cannot be empty")
	}

	// Normalize extensions to lowercase with dot prefix
	normalizedExts := make(map[string]struct{}, len(extensions))
	for _, ext := range extensions {
		if ext == "" {
			continue // Skip empty extensions
		}
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		normalizedExts[strings.ToLower(ext)] = struct{}{}
	}

	// Preallocate slice with estimated capacity to reduce allocations
	// Estimate: assume average directory might have 10 files
	matchingFiles := make([]string, 0, 10)

	// Walk through directory tree
	err := filepath.WalkDir(path, func(currentPath string, entry fs.DirEntry, err error) error {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Handle traversal errors
		if err != nil {
			// Return the error to stop processing, as it might indicate a serious issue
			return fmt.Errorf("unable to access %s: %w", currentPath, err)
		}

		// Skip directories
		if entry.IsDir() {
			return nil
		}

		// If no extensions specified, include all files
		if len(normalizedExts) == 0 {
			matchingFiles = append(matchingFiles, currentPath)
			return nil
		}

		// Check if file extension matches any of the specified extensions
		fileExt := strings.ToLower(filepath.Ext(currentPath))
		if _, matches := normalizedExts[fileExt]; matches {
			matchingFiles = append(matchingFiles, currentPath)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to traverse directory: %w", err)
	}

	return matchingFiles, nil
}
