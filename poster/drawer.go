package poster

import (
	"errors"
	"github.com/golang/freetype/truetype"
	"github.com/minlib/go-util/imagex"
	"image"
	"image/draw"
)

// ##############################
// 核心接口定义
// ##############################

// Drawer 绘制组件接口
type Drawer interface {
	Type() string            // 组件类型标识
	Draw(ctx *Context) error // 执行绘制逻辑
	Validate() error         // 验证配置合法性
}

// ##############################
// 绘制上下文
// ##############################

// Context 绘制上下文，传递全局状态和资源
type Context struct {
	Canvas    draw.Image      // 画布
	Resources *ResourceLoader // 资源加载器
}

// ##############################
// 资源加载器
// ##############################

// ResourceLoader 资源加载器，抽象资源获取逻辑
type ResourceLoader struct{}

// LoadImage 加载图片资源（支持本地路径和远程URL）
func (r *ResourceLoader) LoadImage(path string) (image.Image, error) {
	return imagex.ReadImage(path)
}

// LoadFont 加载字体资源
func (r *ResourceLoader) LoadFont(path string) (*truetype.Font, error) {
	return GetFont(path) // 复用util.go中的字体加载逻辑
}

// ##############################
// 绘制流水线（核心控制器）
// ##############################

// Pipeline 绘制流水线，管理组件执行和输出
type Pipeline struct {
	drawers []Drawer // 有序绘制组件列表
	output  string   // 输出路径（替代FinishDraw）
}

// NewPipeline 创建流水线实例
func NewPipeline(drawers ...Drawer) *Pipeline {
	return &Pipeline{drawers: drawers}
}

// AddDrawer 添加绘制组件
func (p *Pipeline) AddDrawer(drawers ...Drawer) *Pipeline {
	p.drawers = append(p.drawers, drawers...)
	return p
}

// SetOutput 设置输出路径（替代FinishDraw的Output字段）
func (p *Pipeline) SetOutput(output string) *Pipeline {
	p.output = output
	return p
}

// Execute 执行绘制流程
func (p *Pipeline) Execute(ctx *Context) error {
	// 1. 验证所有组件配置
	if err := p.validate(); err != nil {
		return errors.New("pipeline validation failed: " + err.Error())
	}

	// 2. 按顺序执行绘制组件
	for i, drawer := range p.drawers {
		if err := drawer.Draw(ctx); err != nil {
			return errors.New("drawer[" + drawer.Type() + "] failed at index " + string(rune(i)) + ": " + err.Error())
		}
	}

	// 3. 自动保存图片（替代FinishDraw的Do方法）
	if p.output != "" {
		if err := imagex.SavePNG(ctx.Canvas, p.output); err != nil {
			return errors.New("save image failed: " + err.Error())
		}
	}

	return nil
}

// 验证所有组件配置合法性
func (p *Pipeline) validate() error {
	for _, drawer := range p.drawers {
		if err := drawer.Validate(); err != nil {
			return errors.New("invalid drawer " + drawer.Type() + ": " + err.Error())
		}
	}
	return nil
}

// ##############################
// 辅助工具函数
// ##############################

// NewContext 创建绘制上下文
func NewContext(width, height int) *Context {
	return &Context{
		Canvas:    image.NewRGBA(image.Rect(0, 0, width, height)),
		Resources: &ResourceLoader{},
	}
}
