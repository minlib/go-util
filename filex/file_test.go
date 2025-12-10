package filex

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"Basic file", "/home/user/file.txt", "file.txt"},
		{"Directory", "/home/user/", "user"},
		{"Root", "/", "/"},
		{"Empty", "", "."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base(tt.path); got != tt.want {
				t.Errorf("Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"Absolute path", "/home/user/file.txt", "/home/user"},
		{"Relative path", "user/file.txt", "user"},
		{"Directory", "/home/user/", "/home/user"},
		{"Root", "/", "/"},
		{"Empty", "", "."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dir(tt.path); got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExt(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"PDF file", "document.pdf", ".pdf"},
		{"No extension", "README", ""},
		{"Multiple dots", "file.test.go", ".go"},
		{"Hidden file", ".gitignore", ".gitignore"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.filename); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExist(t *testing.T) {
	// Create a temporary file for testing
	tempFile := filepath.Join(os.TempDir(), "test_exist.txt")
	content := "test content"
	err := WriteFileString(tempFile, content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	tests := []struct {
		name string
		path string
		want bool
	}{
		{"Existing file", tempFile, true},
		{"Non-existing file", "/non/existing/file.txt", false},
		{"Empty path", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.path); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := filepath.Join(os.TempDir(), "test_isdir")
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a temporary file for testing
	tempFile := filepath.Join(os.TempDir(), "test_isdir_file.txt")
	err = WriteFileString(tempFile, "test content")
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	tests := []struct {
		name string
		path string
		want bool
	}{
		{"Existing directory", tempDir, true},
		{"Existing file", tempFile, false},
		{"Non-existing path", "/non/existing/path", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.path); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSize(t *testing.T) {
	// Create a temporary file for testing
	tempFile := filepath.Join(os.TempDir(), "test_size.txt")
	content := "Hello, World!"
	err := WriteFileString(tempFile, content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	tests := []struct {
		name     string
		filename string
		want     int64
		wantErr  bool
	}{
		{"Existing file", tempFile, int64(len(content)), false},
		{"Non-existing file", "/non/existing/file.txt", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSize(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		want     string
	}{
		{"Simple path", []string{"/home", "user", "file.txt"}, "/home/user/file.txt"},
		{"With empty elements", []string{"/home", "", "user"}, "/home/user"},
		{"Single element", []string{"file.txt"}, "file.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.elements...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinDir(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		want     string
	}{
		{"Simple path", []string{"/home", "user", "dir"}, "/home/user/dir/"},
		{"With empty elements", []string{"/home", "", "user"}, "/home/user/"},
		{"Single element", []string{"dir"}, "dir/"},
		{"No elements", []string{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinDir(tt.elements...); got != tt.want {
				t.Errorf("JoinDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkdirAll(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "test_mkdirall")
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		paths   []string
		wantErr bool
	}{
		{"Create directories", []string{
			filepath.Join(tempDir, "dir1", "subdir"),
			filepath.Join(tempDir, "dir2", "file.txt"),
		}, false},
		{"Empty paths", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MkdirAll(tt.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkdirAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "test_writefile")
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name     string
		filename string
		data     string
		wantErr  bool
	}{
		{"Normal file", filepath.Join(tempDir, "test.txt"), "test content", false},
		{"File in subdirectory", filepath.Join(tempDir, "subdir", "test.txt"), "test content", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteFileString(tt.filename, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the file was written correctly
			if !tt.wantErr {
				content, err := ReadFileString(tt.filename)
				if err != nil {
					t.Errorf("Failed to read written file: %v", err)
					return
				}
				if content != tt.data {
					t.Errorf("WriteFile() wrote %v, want %v", content, tt.data)
				}
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile := filepath.Join(os.TempDir(), "test_readfile.txt")
	content := "test content for reading"
	err := WriteFileString(tempFile, content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
	}{
		{"Existing file", tempFile, content, false},
		{"Non-existing file", "/non/existing/file.txt", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileString(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFileString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	// Create a temporary source file for testing
	srcFile := filepath.Join(os.TempDir(), "test_copy_src.txt")
	content := "test content for copying"
	err := WriteFileString(srcFile, content)
	if err != nil {
		t.Fatalf("Failed to create source test file: %v", err)
	}
	defer os.Remove(srcFile)

	// Destination file path
	destFile := filepath.Join(os.TempDir(), "test_copy_dest.txt")
	defer os.Remove(destFile)

	tests := []struct {
		name    string
		src     string
		dest    string
		wantErr bool
	}{
		{"Normal copy", srcFile, destFile, false},
		{"Non-existing source", "/non/existing/file.txt", filepath.Join(os.TempDir(), "dest.txt"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CopyFile(tt.src, tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the file was copied correctly
			if !tt.wantErr {
				copiedContent, err := ReadFileString(tt.dest)
				if err != nil {
					t.Errorf("Failed to read copied file: %v", err)
					return
				}
				if copiedContent != content {
					t.Errorf("CopyFile() copied %v, want %v", copiedContent, content)
				}
			}
		})
	}
}

func TestIsExtensions(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		extensions []string
		want       bool
	}{
		{"PDF file with .pdf extension", "document.pdf", []string{".pdf"}, true},
		{"PDF file with pdf extension", "document.pdf", []string{"pdf"}, true},
		{"JPG file case insensitive", "image.JPG", []string{".jpg"}, true},
		{"Multiple extensions match", "document.pdf", []string{".txt", ".pdf"}, true},
		{"No match", "document.txt", []string{".pdf"}, false},
		{"Empty extensions", "document.pdf", []string{}, false},
		{"Empty path", "", []string{".pdf"}, false},
		{"No extension in path", "README", []string{".pdf"}, false},
		{"Empty extension in list", "document.pdf", []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExtensions(tt.path, tt.extensions...); got != tt.want {
				t.Errorf("IsExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByExtensions(t *testing.T) {
	// Create a temporary directory structure for testing
	tempDir := filepath.Join(os.TempDir(), "test_find_extensions")
	defer os.RemoveAll(tempDir)

	// Create test directory structure
	dirs := []string{
		filepath.Join(tempDir, "dir1"),
		filepath.Join(tempDir, "dir2"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}
	}

	// Create test files
	files := []struct {
		path    string
		content string
	}{
		{filepath.Join(tempDir, "file1.txt"), "content1"},
		{filepath.Join(tempDir, "file2.pdf"), "content2"},
		{filepath.Join(tempDir, "dir1", "file3.txt"), "content3"},
		{filepath.Join(tempDir, "dir1", "file4.jpg"), "content4"},
		{filepath.Join(tempDir, "dir2", "file5.pdf"), "content5"},
	}

	for _, f := range files {
		if err := WriteFileString(f.path, f.content); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	tests := []struct {
		name       string
		path       string
		extensions []string
		wantCount  int
		wantErr    bool
	}{
		{"Find all files", tempDir, []string{}, len(files), false},
		{"Find PDF files", tempDir, []string{".pdf"}, 2, false},
		{"Find TXT files", tempDir, []string{".txt"}, 2, false},
		{"Find JPG files", tempDir, []string{".jpg"}, 1, false},
		{"Non-existing directory", "/non/existing/dir", []string{".pdf"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindByExtensions(tt.path, tt.extensions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByExtensions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantCount {
				t.Errorf("FindByExtensions() found %d files; want %d", len(got), tt.wantCount)
			}
		})
	}
}

func TestFindByExtensionsContext(t *testing.T) {
	// Create a temporary directory structure for testing
	tempDir := filepath.Join(os.TempDir(), "test_find_extensions_ctx")
	defer os.RemoveAll(tempDir)

	// Create test files
	files := []struct {
		path    string
		content string
	}{
		{filepath.Join(tempDir, "file1.txt"), "content1"},
		{filepath.Join(tempDir, "file2.pdf"), "content2"},
	}

	for _, f := range files {
		if err := WriteFileString(f.path, f.content); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Test normal operation
	t.Run("Normal operation", func(t *testing.T) {
		ctx := context.Background()
		got, err := FindByExtensionsContext(ctx, tempDir, ".txt")
		if err != nil {
			t.Errorf("FindByExtensionsContext() error = %v", err)
			return
		}
		if len(got) != 1 {
			t.Errorf("FindByExtensionsContext() found %d files; want 1", len(got))
		}
	})

	// Test with cancellation
	t.Run("With cancellation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		// This test might be too fast to actually trigger cancellation
		// but it verifies the function accepts context correctly
		_, err := FindByExtensionsContext(ctx, tempDir, ".txt")
		if err != nil && err != ctx.Err() {
			t.Errorf("FindByExtensionsContext() error = %v", err)
		}
	})
}
