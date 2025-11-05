package drawer

import (
	"fmt"
	"github.com/minlib/go-util/colorx"
	"github.com/minlib/go-util/imagex"
	"image/color"
	"os"
	"testing"
	"time"
)

const fontPath = "../assets/fonts/syht.ttf"
const templatePath = "../outputs/template.png"
const avatarPath = "../outputs/avatar.png"

// getOutputPath generates a unique output path for test images.
func getOutputPath() string {
	return fmt.Sprintf("%s%d.png", "../outputs/", time.Now().UnixNano())
}

// setup is a common method to be executed before each test
func setup() {
	if err := GenerateTestAssets(); err != nil {
		return
	}
	fmt.Println("Test assets generated successfully")
}

// TestMain is the entry point for running tests
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

// GenerateTestAssets creates solid color images for testing
func GenerateTestAssets() error {
	// Create a solid white background template (750x1334)
	template := imagex.NewImage(750, 1334, color.RGBA{R: 235, G: 235, B: 235, A: 255})
	if err := imagex.SavePNG(template, templatePath); err != nil {
		return err
	}
	// Create a solid blue avatar (200x200)
	avatar := imagex.NewImage(200, 200, color.RGBA{R: 65, G: 105, B: 225, A: 255})
	if err := imagex.SavePNG(avatar, avatarPath); err != nil {
		return err
	}
	return nil
}

// TestPipelineExample demonstrates how to use the new Pipeline API to create a image.
func TestPipelineExample(t *testing.T) {
	canvas := imagex.NewImage(750, 1334, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	// Create drawing context
	ctx := NewContext(canvas)
	// Create drawing components
	avatarDraw := &ImageDraw{
		X:     30,
		Y:     50,
		Path:  avatarPath,
		Round: true,
	}
	textDraw := &TextDraw{
		Left:     180,
		Top:      150,
		Size:     26,
		Color:    "#000000",
		Content:  "这里是标题文字",
		FontPath: fontPath,
	}
	qrCodeDraw := &QRCodeDraw{
		X:       30,
		Y:       860,
		Size:    250,
		Content: "https://www.minzhan.com",
	}
	// Create pipeline and add components
	outputPath := getOutputPath()
	pipeline := NewPipeline().
		AddDrawer(avatarDraw).
		AddDrawer(textDraw).
		AddDrawer(qrCodeDraw).
		SetOutput(outputPath)

	// Execute drawing pipeline
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}

	fmt.Printf("Pipeline execution success, output path: %s\n", outputPath)
}

// TestBuilderExample demonstrates how to use Builder to create a poster.
func TestBuilderExample(t *testing.T) {
	canvas := imagex.NewImage(750, 1334, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	// Use Builder to create canvas
	builder := NewBuilder(canvas)
	// Add drawing components
	builder.AddDrawer(
		&ImageDraw{
			X:     30,
			Y:     50,
			Path:  avatarPath,
			Round: true,
		},
		&TextDraw{
			Left:     180,
			Top:      105,
			Size:     26,
			Color:    "#FFFFFF",
			Content:  "这里是大标题1",
			FontPath: fontPath,
		},
		&TextDraw{
			Left:     180,
			Top:      150,
			Size:     20,
			Color:    "#FFFFFF",
			Content:  "这里是小标题2",
			FontPath: fontPath,
		},
		&QRCodeDraw{
			X:       30,
			Y:       860,
			Size:    250,
			Content: "https://www.minzhan.com",
		},
	)
	// Build and save image
	image, err := builder.Build()
	if err != nil {
		t.Errorf("Build image failed: %v", err)
		return
	}
	// Save image
	outputPath := getOutputPath()
	if err = imagex.SavePNG(image, outputPath); err != nil {
		t.Errorf("Save image failed: %v", err)
		return
	}
	fmt.Printf("Builder success, output path: %s\n", outputPath)
}

// TestTextAlignment demonstrates text drawing with different alignments.
func TestTextAlignment(t *testing.T) {
	canvas := imagex.NewImage(750, 1334, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	width := canvas.Bounds().Dx()
	ctx := NewContext(canvas)
	// Create drawing components
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: templatePath,
	}
	// Draw text
	textDraw1 := &TextDraw{
		Left:     132,
		Top:      190,
		Size:     28,
		Color:    "#A5A6A8",
		Content:  "Minzhan All Rights Reserved.",
		FontPath: fontPath,
	}
	textDraw2 := &TextDraw{
		Left:     132,
		Top:      236,
		Size:     30,
		Color:    "#A5A6A8",
		Content:  "民站科技（深圳）有限公司",
		FontPath: fontPath,
	}
	// Single line text left alignment
	textDraw3 := &TextDraw{
		Left:     20,
		Top:      300,
		AX:       FlexStart,
		AY:       FlexCenter,
		Size:     30,
		Color:    "#999999",
		Content:  "居左对齐",
		FontPath: fontPath,
	}
	// Single line text center alignment
	textDraw4 := &TextDraw{
		Left:        float64(width / 2),
		Top:         330,
		AX:          FlexCenter,
		AY:          FlexCenter,
		Size:        30,
		Color:       "#999999",
		Content:     "居中对齐居中对齐居中对齐",
		FontPath:    fontPath,
		CorrectionY: -5,
	}
	// Single line text right alignment
	textDraw5 := &TextDraw{
		Left:        float64(width - 20),
		Top:         360,
		AX:          FlexEnd,
		AY:          FlexCenter,
		Size:        30,
		Color:       "#999999",
		Content:     "居右对齐",
		FontPath:    fontPath,
		CorrectionY: -5,
	}
	// Multi-line text left alignment
	textDraw6 := &TextDraw{
		Left:        20,
		Top:         400,
		AX:          FlexStart,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居左对齐\n第二行文本\n第三行文本",
		FontPath:    fontPath,
		Align:       AlignLeft,
		LineSpacing: 1,
		CorrectionY: -5,
	}
	// Multi-line text center alignment
	textDraw7 := &TextDraw{
		Left:        float64(width / 2),
		Top:         550,
		AX:          FlexCenter,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居中对齐\n第二行文本\n第三行文本",
		FontPath:    fontPath,
		Align:       AlignCenter,
		LineSpacing: 1.5,
		CorrectionY: -5,
	}
	// Multi-line text right alignment
	textDraw8 := &TextDraw{
		Left:        float64(width) - 20,
		Top:         700,
		AX:          FlexEnd,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居右对齐\n第二行文本\n第三行文本",
		FontPath:    fontPath,
		Align:       AlignRight,
		LineSpacing: 1.5,
		CorrectionY: -5,
	}
	// Multi-line text left alignment with custom line spacing
	textDraw9 := &TextDraw{
		Left:        20,
		Top:         850,
		Size:        12,
		Color:       "#000000",
		Content:     "APP开发需要注意的细节非常多，这里罗列一些，避免大家踩坑：\n1、找开发公司或者APP开发团队。要多渠道的找，找一些觉得靠谱的开发公司，多接触，创始人最好是技术出身，有技术基因的公司才能保证项目的开发质量。很多的做销售出身的公司，技术能力真心没有办法保证。\n2、合同签订。合同内容一定要细致，需要有比较详细的列表和功能描述，这样才能保证后期不会出现扯皮。因为软件开发需求经常会变，开发公司有时候也有愉懒的情况。\n3、需求沟通。沟通结果一定要落实到纸或者邮件、文档。最后要产生详尽的产品原型。原型是必须的，产品文档可根据实际情况来确定要不要，因为产品文档这个太需要时间，可能咱的费用及开发公司精力等方面限制，PRD文档不是必须的。\n4、产品研发。一定要提前沟通好技术架构，这样对项目开发内沟通，以及后续产品版本迭代都会有非常大的帮助，减少沟通成本，提高开发效率和质量。\n5、产品测试。这个环节非常重要，咱们需要在beta版本的时候参与进来。这样可以更早的了解熟悉软件的实现情况，为后续运营作好准备。\n6、产品验收。一定要把各个功能细节，一定都要过2~5遍。这样双方都放心一些。 ",
		FontPath:    fontPath,
		Align:       AlignLeft,
		LineSpacing: 1,
	}
	// Create pipeline
	outputPath := getOutputPath()
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		AddDrawer(textDraw1).
		AddDrawer(textDraw2).
		AddDrawer(textDraw3).
		AddDrawer(textDraw4).
		AddDrawer(textDraw5).
		AddDrawer(textDraw6).
		AddDrawer(textDraw7).
		AddDrawer(textDraw8).
		AddDrawer(textDraw9).
		SetOutput(outputPath)
	// Execute drawing
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Text test failed: %v", err)
		return
	}
	fmt.Printf("Text test success, output path: %s\n", outputPath)
}

// TestCircle demonstrates how to draw circular images.
func TestCircle(t *testing.T) {
	srcImage, err := imagex.ReadImage(avatarPath)
	if err != nil {
		t.Skipf("Skip test due to missing avatar: %v", err)
		return
	}
	newImage := DrawCircle(srcImage)
	outputPath := getOutputPath()
	if err = imagex.SavePNG(newImage, outputPath); err != nil {
		t.Errorf("Failed to save circle image: %v", err)
		return
	}
	fmt.Printf("Circle test success, output path: %s\n", outputPath)
}

// TestProductImage demonstrates how to draw product posters.
func TestProductImage(t *testing.T) {
	canvas := imagex.NewImage(750, 1100, colorx.Hex2RGBA("#FFFFFFFF"))
	ctx := NewContext(canvas)
	avatarDraw1 := &ImageDraw{
		X:     30,
		Y:     30,
		Path:  "https://static.minzhan.com/uploads/s274600676091367425/thumb/202406/1331ff9d1377f549450280b7509786308ad3.webp",
		Width: 690,
		Round: false,
	}
	avatarDraw2 := &ImageDraw{
		X:     500,
		Y:     820,
		Path:  "https://res.wx.qq.com/wxdoc/dist/assets/img/skyline-demo.37eff20b.png",
		Width: 200,
		Round: true,
	}
	textDraw1 := &TextDraw{
		Left:        float64(750) - 30,
		Top:         1040,
		AX:          FlexEnd,
		Size:        28,
		Color:       "#000000",
		Content:     "长按识别小程序码",
		FontPath:    fontPath,
		LineSpacing: 1,
		Align:       AlignRight,
	}
	textDraw2 := &TextDraw{
		Left:        30,
		Top:         750,
		Size:        36,
		Color:       "#000000",
		Content:     "马面裙白色上衣红色织金妆花\n双工艺面料马面裙",
		FontPath:    fontPath,
		LineSpacing: 1.2,
		Align:       AlignLeft,
	}
	textDraw3 := &TextDraw{
		Left:     30,
		Top:      860,
		Size:     40,
		Color:    "#FF0000",
		Content:  "100元",
		FontPath: fontPath,
	}
	// Create pipeline and add components
	outputPath := getOutputPath()
	pipeline := NewPipeline().
		AddDrawer(avatarDraw1).
		AddDrawer(avatarDraw2).
		AddDrawer(textDraw1).
		AddDrawer(textDraw2).
		AddDrawer(textDraw3).
		SetOutput(outputPath)
	// Execute drawing pipeline
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}
	fmt.Printf("Product image test success, output path: %s\n", outputPath)
}
