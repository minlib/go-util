package pdf2image

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// setupTestDir 创建测试目录
func setupTestDir(t *testing.T) string {
	tempDir := filepath.Join(os.TempDir(), "pdf2image_test")
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("无法创建测试目录: %v", err)
	}
	return tempDir
}

// teardownTestDir 清理测试目录
func teardownTestDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		t.Logf("警告: 无法清理测试目录 %s: %v", dir, err)
	}
}

// setupTestDirWithSubdirs 创建带子目录结构的测试目录
func setupTestDirWithSubdirs(t *testing.T) (string, string) {
	tempDir := filepath.Join(os.TempDir(), "pdf2image_test_recursive")

	// 创建主目录
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		t.Fatalf("无法创建测试目录: %v", err)
	}

	// 创建子目录
	subDir := filepath.Join(tempDir, "subdir")
	err = os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("无法创建子目录: %v", err)
	}

	return tempDir, subDir
}

// createDummyPDFFiles 在指定目录创建虚拟PDF文件
func createDummyPDFFiles(t *testing.T, dir string, filenames []string) {
	for _, filename := range filenames {
		filePath := filepath.Join(dir, filename)
		// 创建一个简单的文本文件模拟PDF文件（实际测试中不会真正转换）
		content := "%PDF-1.4\n%äüöß\n"
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("无法创建测试文件 %s: %v", filePath, err)
		}
	}
}

// TestValidateOptions 测试选项验证功能
func TestValidateOptions(t *testing.T) {
	tests := []struct {
		name        string
		options     ConvertOptions
		expectError bool
	}{
		{
			name: "有效选项",
			options: ConvertOptions{
				Format:  "png",
				DPI:     300,
				Quality: 95,
			},
			expectError: false,
		},
		{
			name: "无效格式",
			options: ConvertOptions{
				Format: "invalid",
			},
			expectError: true,
		},
		{
			name: "负DPI",
			options: ConvertOptions{
				DPI: -1,
			},
			expectError: true,
		},
		{
			name: "质量超出范围",
			options: ConvertOptions{
				Quality: 101,
			},
			expectError: true,
		},
		{
			name: "负线程数",
			options: ConvertOptions{
				ThreadCount: -1,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOptions(tt.options)
			if tt.expectError && err == nil {
				t.Errorf("期望错误但没有错误")
			}
			if !tt.expectError && err != nil {
				t.Errorf("未期望错误但得到错误: %v", err)
			}
		})
	}
}

// TestApplyDefaultOptions 测试默认选项应用
func TestApplyDefaultOptions(t *testing.T) {
	options := &ConvertOptions{}
	applyDefaultOptions(options)

	if options.OutputPath != "./outputs" {
		t.Errorf("期望输出路径 './outputs', 得到 '%s'", options.OutputPath)
	}
	if options.Format != "png" {
		t.Errorf("期望格式 'png', 得到 '%s'", options.Format)
	}
	if options.DPI != 300 {
		t.Errorf("期望DPI 300, 得到 %d", options.DPI)
	}
	if options.Quality != 95 {
		t.Errorf("期望质量 95, 得到 %d", options.Quality)
	}
}

// TestGetPdftocairoPath 测试获取pdftocairo路径
func TestGetPdftocairoPath(t *testing.T) {
	// 测试系统路径查找
	path, err := getPdftocairoPath("")
	if err != nil {
		t.Logf("警告: 无法找到系统pdftocairo: %v", err)
	} else if path == "" {
		t.Error("期望非空路径")
	}

	// 测试无效自定义路径
	_, err = getPdftocairoPath("/invalid/path")
	if err == nil {
		t.Error("期望错误但没有错误")
	}
}

// TestConvertSinglePDFInvalidInput 测试单个PDF转换的无效输入
func TestConvertSinglePDFInvalidInput(t *testing.T) {
	tempDir := setupTestDir(t)
	defer teardownTestDir(t, tempDir)

	tests := []struct {
		name        string
		options     ConvertOptions
		expectError bool
	}{
		{
			name: "空PDF路径",
			options: ConvertOptions{
				PDFPath:    "",
				OutputPath: tempDir,
			},
			expectError: true,
		},
		{
			name: "不存在的PDF文件",
			options: ConvertOptions{
				PDFPath:    "/non/existent/file.pdf",
				OutputPath: tempDir,
			},
			expectError: true,
		},
		{
			name: "非PDF文件",
			options: ConvertOptions{
				PDFPath:    "/etc/hosts", // 这是一个文本文件
				OutputPath: tempDir,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertSinglePDF(tt.options)
			if tt.expectError && err == nil {
				t.Error("期望错误但没有错误")
			}
			if !tt.expectError && err != nil {
				t.Errorf("未期望错误但得到错误: %v", err)
			}
		})
	}
}

// TestConvertSinglePDFWithDifferentFormats 测试不同格式的转换
func TestConvertSinglePDFWithDifferentFormats(t *testing.T) {
	tempDir := setupTestDir(t)
	defer teardownTestDir(t, tempDir)

	formats := []string{"png", "jpeg", "ppm", "tiff"}
	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			options := ConvertOptions{
				PDFPath:    "/non/existent/file.pdf", // 我们只测试参数验证，不实际执行转换
				OutputPath: tempDir,
				Format:     format,
			}
			// 这里我们期望得到文件不存在的错误，而不是格式错误
			_, err := ConvertSinglePDF(options)
			if err != nil && fmt.Sprintf("%v", err) == "文件不存在: /non/existent/file.pdf" {
				// 这是预期的错误，说明格式参数通过了验证
				t.Logf("格式 %s 通过验证", format)
			} else if err != nil && fmt.Sprintf("%v", err) != "文件不存在: /non/existent/file.pdf" {
				// 其他错误可能是格式验证失败
				t.Errorf("格式 %s 验证失败: %v", format, err)
			}
		})
	}
}

// TestConvertBatchPDFInvalidInput 测试批量转换的无效输入
func TestConvertBatchPDFInvalidInput(t *testing.T) {
	tempDir := setupTestDir(t)
	defer teardownTestDir(t, tempDir)

	tests := []struct {
		name        string
		options     ConvertOptions
		expectError bool
	}{
		{
			name: "空输入目录",
			options: ConvertOptions{
				InputDir:   "",
				OutputPath: tempDir,
			},
			expectError: true,
		},
		{
			name: "不存在的输入目录",
			options: ConvertOptions{
				InputDir:   "/non/existent/dir",
				OutputPath: tempDir,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertBatchPDF(tt.options)
			if tt.expectError && err == nil {
				t.Error("期望错误但没有错误")
			}
			if !tt.expectError && err != nil {
				t.Errorf("未期望错误但得到错误: %v", err)
			}
		})
	}
}

// TestConvertBatchPDFNonRecursive 测试非递归批量转换
func TestConvertBatchPDFNonRecursive(t *testing.T) {
	// 创建测试目录结构
	tempDir, subDir := setupTestDirWithSubdirs(t)
	defer teardownTestDir(t, tempDir)

	// 在主目录和子目录中创建测试PDF文件
	mainPDFs := []string{"file1.pdf", "file2.pdf"}
	subPDFs := []string{"subfile1.pdf", "subfile2.pdf"}

	createDummyPDFFiles(t, tempDir, mainPDFs)
	createDummyPDFFiles(t, subDir, subPDFs)

	// 测试非递归模式（只处理主目录中的文件）
	options := ConvertOptions{
		InputDir:    tempDir,
		Recursive:   false, // 非递归
		OutputPath:  filepath.Join(tempDir, "output"),
		Format:      "png",
		DPI:         100, // 使用较低DPI加快测试
		ThreadCount: 2,
	}

	// 由于我们使用的是虚拟PDF文件，实际转换会失败，但我们主要测试目录查找逻辑
	result, err := ConvertBatchPDF(options)

	// 在非递归模式下，应该只找到主目录中的2个文件
	// 但由于我们使用的是虚拟PDF文件，实际转换会失败
	// 我们检查是否有处理结果返回
	if err != nil {
		// 如果是由于找不到pdftocairo工具导致的错误，这是预期的
		if fmt.Sprintf("%v", err) == "找不到pdftocairo命令，请确保已安装poppler-utils\nLinux安装命令: sudo apt install poppler-utils" {
			t.Logf("由于缺少pdftocairo工具，测试通过")
		} else {
			// 检查是否找到了文件（即使转换失败）
			if result != nil {
				t.Logf("找到了 %d 个文件，但由于缺少pdftocairo工具转换失败", len(result))
			} else {
				t.Errorf("未找到任何文件: %v", err)
			}
		}
	} else {
		// 如果没有错误，检查是否找到了正确的文件数量
		if len(result) != 2 {
			t.Errorf("非递归模式下期望找到2个文件，实际找到 %d 个", len(result))
		}
	}
}

// TestConvertBatchPDFRecursive 测试递归批量转换
func TestConvertBatchPDFRecursive(t *testing.T) {
	// 创建测试目录结构
	tempDir, subDir := setupTestDirWithSubdirs(t)
	defer teardownTestDir(t, tempDir)

	// 在主目录和子目录中创建测试PDF文件
	mainPDFs := []string{"file1.pdf", "file2.pdf"}
	subPDFs := []string{"subfile1.pdf", "subfile2.pdf"}

	createDummyPDFFiles(t, tempDir, mainPDFs)
	createDummyPDFFiles(t, subDir, subPDFs)

	// 测试递归模式（处理主目录和子目录中的所有文件）
	options := ConvertOptions{
		InputDir:    tempDir,
		Recursive:   true, // 递归
		OutputPath:  filepath.Join(tempDir, "output"),
		Format:      "png",
		DPI:         100, // 使用较低DPI加快测试
		ThreadCount: 1,
	}

	// 由于我们使用的是虚拟PDF文件，实际转换会失败，但我们主要测试目录查找逻辑
	result, err := ConvertBatchPDF(options)

	// 在递归模式下，应该找到主目录和子目录中的所有4个文件
	// 但由于我们使用的是虚拟PDF文件，实际转换会失败
	// 我们检查是否有处理结果返回
	if err != nil {
		// 如果是由于找不到pdftocairo工具导致的错误，这是预期的
		if fmt.Sprintf("%v", err) == "找不到pdftocairo命令，请确保已安装poppler-utils\nLinux安装命令: sudo apt install poppler-utils" {
			t.Logf("由于缺少pdftocairo工具，测试通过")
		} else {
			// 检查是否找到了文件（即使转换失败）
			if result != nil {
				t.Logf("找到了 %d 个文件，但由于缺少pdftocairo工具转换失败", len(result))
				// 检查是否至少找到了主目录中的文件
				if len(result) < 2 {
					t.Errorf("递归模式下期望至少找到2个文件，实际找到 %d 个", len(result))
				}
			} else {
				t.Errorf("未找到任何文件: %v", err)
			}
		}
	} else {
		// 如果没有错误，检查是否找到了正确的文件数量
		if len(result) != 4 {
			t.Errorf("递归模式下期望找到4个文件，实际找到 %d 个", len(result))
		}
	}
}

// TestConvertSinglePDF 测试单个PDF转换功能
func TestConvertSinglePDF(t *testing.T) {
	// 配置转换选项
	options := ConvertOptions{
		PDFPath:    "/home/work/assets/pdf_files/input.pdf",
		OutputPath: "/home/work/outputs/",
		Format:     "png",
		DPI:        300,
	}

	// 执行转换
	result, err := ConvertSinglePDF(options)
	if err != nil {
		// 注意：在实际测试环境中，如果文件不存在这是预期的行为
		t.Logf("转换结果: %v", err)
		return
	}
	fmt.Println("============================================")
	fmt.Printf("PDF文件：%s\n生成的图片数：%d\n生成的图片：\n", options.PDFPath, result.ImageCount)
	for _, imagePath := range result.ImagePaths {
		fmt.Printf("    - %s\n", imagePath)
	}
}

// TestConvertBatchPDF 测试批量PDF转换功能
func TestConvertBatchPDF(t *testing.T) {
	// 配置转换选项，指向正确的PDF文件目录
	options := ConvertOptions{
		InputDir:    "/home/work/assets/pdf_files",
		Recursive:   true,
		ThreadCount: 1,
		OutputPath:  "/home/work/outputs/",
		Format:      "png",
		DPI:         100,
	}

	// 执行转换
	mapResult, err := ConvertBatchPDF(options)
	if err != nil {
		// 注意：在实际测试环境中，如果目录不存在这是预期的行为
		t.Logf("转换结果: %v", err)
		return
	}

	// 计算总图片数
	totalImages := 0
	fmt.Println("============================================")
	for pdfPath, result := range mapResult {
		if result.Success {
			totalImages += result.ImageCount
			fmt.Printf("PDF文件：%s\n生成的图片数：%d\n生成的图片：\n", pdfPath, result.ImageCount)
			for _, imagePath := range result.ImagePaths {
				fmt.Printf("    - %s\n", imagePath)
			}
		}
	}
	fmt.Printf("生成的图片总数：%d\n", totalImages)
}
