// Package pdf2image provides functionality to convert PDF files to images.
//
// Requirements:
// This package requires the pdftocairo tool from Poppler-utils to be installed on the system.
//
// Installation of Poppler-utils:
//   - Ubuntu/Debian: sudo apt-get install poppler-utils
//   - CentOS/RHEL: sudo yum install poppler-utils
//   - Fedora: sudo dnf install poppler-utils
//   - macOS: brew install poppler
//   - Windows: Download from https://github.com/oschwartz1066/poppler-windows/releases
//
// Usage:
//   - Single PDF conversion: ConvertSinglePDF(ConvertOptions)
//   - Batch PDF conversion: ConvertBatchPDF(ConvertOptions)
package pdf2image

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// ConvertOptions 转换选项（合并了单个文件和批量转换的所有参数）
type ConvertOptions struct {
	OutputPath  string // 输出目录
	Format      string // 输出格式 (png/jpeg/ppm/tiff)
	DPI         int    // 分辨率
	Quality     int    // 图片质量 (仅JPEG有效)
	PopplerPath string // Poppler路径
	ThreadCount int    // 线程数
	PDFPath     string // 单个PDF文件路径（用于单文件转换）
	StartPage   int    // 起始页码，单文件转换参数 (0表示第一页)
	EndPage     int    // 结束页码，单文件转换参数 (0表示最后一页)
	InputDir    string // 输入目录（用于批量转换）
	Recursive   bool   // 是否递归处理子目录
}

// ConvertResult 单个PDF转换结果
type ConvertResult struct {
	Success    bool     `json:"success"`
	ImageCount int      `json:"imageCount"`
	ImagePaths []string `json:"imagePaths,omitempty"`
	Error      string   `json:"error,omitempty"`
}

// applyDefaultOptions 应用默认选项
func applyDefaultOptions(options *ConvertOptions) {
	if options.OutputPath == "" {
		options.OutputPath = "./outputs"
	}
	if options.Format == "" {
		options.Format = "png"
	}
	if options.DPI == 0 {
		options.DPI = 300
	}
	if options.Quality == 0 {
		options.Quality = 95
	}
}

// validateOptions 验证转换选项
func validateOptions(options ConvertOptions) error {
	// 验证输出格式
	validFormats := map[string]bool{
		"png":  true,
		"jpeg": true,
		"ppm":  true,
		"tiff": true,
	}
	if options.Format != "" && !validFormats[options.Format] {
		return fmt.Errorf("不支持的输出格式: %s", options.Format)
	}

	// 验证DPI
	if options.DPI < 0 {
		return fmt.Errorf("DPI不能为负数")
	}

	// 验证图片质量
	if options.Quality < 0 || options.Quality > 100 {
		return fmt.Errorf("图片质量必须在0-100之间")
	}

	// 验证线程数
	if options.ThreadCount < 0 {
		return fmt.Errorf("线程数不能为负数")
	}

	return nil
}

// getPdftocairoPath 获取pdftocairo工具路径
func getPdftocairoPath(popplerPath string) (string, error) {
	// 如果设置了自定义路径，则使用自定义路径
	if popplerPath != "" {
		customPath := filepath.Join(popplerPath, "pdftocairo")
		if _, err := os.Stat(customPath); err == nil {
			return customPath, nil
		}
		return "", fmt.Errorf("找不到指定路径的pdftocairo: %s", customPath)
	}

	// 否则在系统PATH中查找
	path, err := exec.LookPath("pdftocairo")
	if err != nil {
		return "", fmt.Errorf("找不到pdftocairo命令，请确保已安装poppler-utils\nLinux安装命令: sudo apt install poppler-utils")
	}
	return path, nil
}

// ConvertSinglePDF 将单个PDF文件转换为图片
func ConvertSinglePDF(options ConvertOptions) (ConvertResult, error) {
	applyDefaultOptions(&options)

	result := ConvertResult{
		Success: false,
	}
	// 验证选项
	if err := validateOptions(options); err != nil {
		result.Error = err.Error()
		return result, err
	}

	pdfPath := options.PDFPath
	startPage := options.StartPage
	endPage := options.EndPage

	// 参数验证
	if pdfPath == "" {
		err := fmt.Errorf("PDF文件路径不能为空")
		result.Error = err.Error()
		return result, err
	}
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		err = fmt.Errorf("文件不存在: %s", pdfPath)
		result.Error = err.Error()
		return result, err
	}
	if !strings.HasSuffix(strings.ToLower(pdfPath), ".pdf") {
		err := fmt.Errorf("不是PDF文件: %s", pdfPath)
		result.Error = err.Error()
		return result, err
	}

	if err := os.MkdirAll(options.OutputPath, 0755); err != nil {
		err = fmt.Errorf("创建输出目录失败: %v\n建议检查目录权限", err)
		result.Error = err.Error()
		return result, err
	}

	pdftocairoPath, err := getPdftocairoPath(options.PopplerPath)
	if err != nil {
		result.Error = err.Error()
		return result, err
	}

	// 核心修复：参数拆分为独立元素（每个参数单独作为切片元素）
	args := []string{
		"-" + options.Format,            // 输出格式（如 -png）
		"-r", strconv.Itoa(options.DPI), // 分辨率（拆分为 "-r" 和 "300"）
	}

	// JPEG质量参数（同样拆分）
	if options.Format == "jpeg" {
		args = append(args, "-quality", strconv.Itoa(options.Quality))
	}

	// 页码范围（拆分参数）
	if startPage > 0 {
		args = append(args, "-f", strconv.Itoa(startPage))
	}
	if endPage > 0 && endPage >= startPage {
		args = append(args, "-l", strconv.Itoa(endPage))
	}

	// 处理输出文件名，直接使用UUID确保唯一性
	filename := strings.TrimSuffix(filepath.Base(pdfPath), ".pdf")

	// 生成UUID以确保唯一性
	uniquePrefix := fmt.Sprintf("%s_%s", filename, uuid.New().String()[:8])
	outputPrefix := filepath.Join(options.OutputPath, fmt.Sprintf("%s_page", uniquePrefix))
	args = append(args, pdfPath, outputPrefix)

	// 执行命令
	fmt.Printf("转换中: %s\n", filepath.Base(pdfPath))
	cmd := exec.Command(pdftocairoPath, args...) // 正确传递参数切片
	output, err := cmd.CombinedOutput()
	if err != nil {
		errorMsg := fmt.Sprintf("转换失败: %v\n命令: %s %s\n输出: %s",
			err, pdftocairoPath, strings.Join(args, " "), string(output))

		// Linux特有错误提示
		switch {
		case strings.Contains(string(output), "unrecognized option"):
			errorMsg += "\n可能原因：参数格式错误（已修复，请重新尝试）"
		case strings.Contains(string(output), "No such file"):
			errorMsg += "\n解决方案：安装poppler-utils (sudo apt install poppler-utils)"
		case strings.Contains(string(output), "Permission denied"):
			errorMsg += "\n解决方案：添加权限 (chmod -R 755 " + options.OutputPath + ")"
		case strings.Contains(string(output), "Failed to open PDF"):
			errorMsg += "\n解决方案：检查PDF是否损坏或加密"
		case strings.Contains(string(output), "Syntax Error"):
			errorMsg += "\n可能原因：PDF文件格式错误或已损坏"
		}

		result.Error = errorMsg
		return result, fmt.Errorf(errorMsg)
	}

	// 统计生成的图片
	imageExt := options.Format
	if imageExt == "jpeg" {
		imageExt = "jpg"
	}

	// 使用filepath.Glob查找生成的图片
	imagePattern := fmt.Sprintf("%s*.%s", outputPrefix, imageExt)
	images, err := filepath.Glob(imagePattern)
	if err != nil {
		err = fmt.Errorf("统计图片失败: %v", err)
		result.Error = err.Error()
		return result, err
	}

	if len(images) == 0 {
		err = fmt.Errorf("未生成图片，可能PDF为空或参数错误")
		result.Error = err.Error()
		return result, err
	}

	// 按文件名排序图片路径
	sort.Strings(images)

	// 设置结果
	result.Success = true
	result.ImageCount = len(images)
	result.ImagePaths = images

	return result, nil
}

// ConvertBatchPDF 批量转换PDF
func ConvertBatchPDF(options ConvertOptions) (map[string]ConvertResult, error) {
	applyDefaultOptions(&options)
	if options.ThreadCount == 0 {
		options.ThreadCount = 4
	}

	// 验证选项
	if err := validateOptions(options); err != nil {
		return nil, err
	}

	pdfDir := options.InputDir
	recursive := options.Recursive

	absDir, err := filepath.Abs(pdfDir)
	if err != nil {
		return nil, fmt.Errorf("解析目录路径失败: %v", err)
	}
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", absDir)
	}

	// 改进的文件查找逻辑
	var pdfFiles []string

	if recursive {
		// 使用filepath.WalkDir进行递归查找
		err := filepath.WalkDir(absDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// 只处理PDF文件，跳过目录
			if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".pdf") {
				pdfFiles = append(pdfFiles, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("遍历目录失败: %v", err)
		}
	} else {
		// 非递归模式，只查找当前目录的PDF文件
		searchPattern := filepath.Join(absDir, "*.pdf")
		pdfFiles, err = filepath.Glob(searchPattern)
		if err != nil {
			return nil, fmt.Errorf("查找PDF失败: %v", err)
		}
	}

	if len(pdfFiles) == 0 {
		return nil, fmt.Errorf("未找到PDF文件: %s", absDir)
	}

	fmt.Printf("找到 %d 个PDF，使用 %d 线程处理...\n", len(pdfFiles), options.ThreadCount)

	var wg sync.WaitGroup
	result := make(map[string]ConvertResult)
	mu := &sync.Mutex{}
	sem := make(chan struct{}, options.ThreadCount)

	for _, pdfPath := range pdfFiles {
		wg.Add(1)
		sem <- struct{}{}

		go func(path string) {
			defer wg.Done()
			defer func() { <-sem }()

			// 为每个文件创建临时选项
			singleOptions := ConvertOptions{
				PDFPath:     path,
				InputDir:    absDir, // 添加InputDir用于生成唯一文件名
				OutputPath:  options.OutputPath,
				Format:      options.Format,
				DPI:         options.DPI,
				Quality:     options.Quality,
				PopplerPath: options.PopplerPath,
			}

			convertResult, err := ConvertSinglePDF(singleOptions)
			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				result[path] = convertResult
				// fmt.Printf("失败: %s - %s\n", filepath.Base(path), err.Error())
			} else {
				result[path] = convertResult
				// fmt.Printf("成功: %s（%d张图）\n", filepath.Base(path), convertResult.ImageCount)
				// fmt.Printf("生成的图片:\n")
				// for _, imagePath := range convertResult.ImagePaths {
				// 	fmt.Printf("    - %s\n", filepath.Base(imagePath))
				// }
			}
		}(pdfPath)
	}

	wg.Wait()
	return result, nil
}

// parseFlags 解析命令行参数
func parseFlags() (ConvertOptions, error) {
	var options ConvertOptions

	// 按 ConvertOptions 字段定义顺序设置 flag
	flag.StringVar(&options.OutputPath, "output", "./outputs", "输出目录")
	flag.StringVar(&options.Format, "format", "png", "输出格式（png/jpeg/ppm/tiff）")
	flag.IntVar(&options.DPI, "dpi", 300, "分辨率")
	flag.IntVar(&options.Quality, "quality", 95, "图片质量（0-100，仅jpeg有效）")
	flag.StringVar(&options.PopplerPath, "poppler", "", "poppler路径（如/usr/bin）")
	flag.IntVar(&options.ThreadCount, "threads", 4, "线程数")
	flag.StringVar(&options.PDFPath, "single", "", "单个PDF文件路径（与--batch二选一）")
	flag.IntVar(&options.StartPage, "start", 0, "起始页码（0表示第一页）")
	flag.IntVar(&options.EndPage, "end", 0, "结束页码（0表示最后一页）")
	flag.StringVar(&options.InputDir, "batch", "", "批量转换目录（与--single二选一）")
	flag.BoolVar(&options.Recursive, "recursive", false, "批量转换时递归子目录")

	flag.Parse()

	if (options.PDFPath == "" && options.InputDir == "") || (options.PDFPath != "" && options.InputDir != "") {
		return options, fmt.Errorf("必须指定--single或--batch，且只能选一个")
	}

	if options.PDFPath != "" && options.StartPage > 0 && options.EndPage > 0 && options.StartPage > options.EndPage {
		return options, fmt.Errorf("起始页码不能大于结束页码: %d > %d", options.StartPage, options.EndPage)
	}

	return options, nil
}

// Command 根据命令行参数执行转换任务
func Command() error {
	options, err := parseFlags()
	if err != nil {
		return fmt.Errorf("参数错误: %v\n使用帮助: --help", err)
	}

	// 根据参数选择执行单个文件转换还是批量转换
	if options.PDFPath != "" {
		result, err := ConvertSinglePDF(options)
		if err != nil {
			return fmt.Errorf("转换失败: %v", err)
		}
		fmt.Printf("\n成功生成 %d 张图片，输出目录: %s\n", result.ImageCount, options.OutputPath)
		fmt.Printf("生成的图片:\n")
		for _, imagePath := range result.ImagePaths {
			fmt.Printf("  - %s\n", filepath.Base(imagePath))
		}
	} else if options.InputDir != "" {
		result, err := ConvertBatchPDF(options)
		if err != nil {
			return fmt.Errorf("批量转换失败: %v", err)
		}

		success, fail := 0, 0
		totalImages := 0
		for _, v := range result {
			if v.Success {
				success++
				totalImages += v.ImageCount
			} else {
				fail++
			}
		}
		fmt.Printf("\n批量转换完成: 成功 %d 个, 失败 %d 个\n输出目录: %s\n", success, fail, options.OutputPath)
		fmt.Printf("总共生成 %d 张图片\n", totalImages)
	}

	return nil
}

func main() {
	if err := Command(); err != nil {
		log.Fatal(err)
	}
}
