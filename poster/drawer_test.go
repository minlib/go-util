package poster

import (
	"fmt"
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
	"testing"
	"time"
)

// TestPipelineExample 展示如何使用新的Pipeline API
func TestPipelineExample(t *testing.T) {
	// 创建绘制上下文
	ctx := NewContext(750, 1334)

	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}

	avatarDraw := &ImageDraw{
		X:     30,
		Y:     50,
		Path:  "./assets/avatar.jpg",
		Round: true,
	}

	textDraw := &TextDraw{
		X:        180,
		Y:        150,
		Size:     26,
		Color:    "#FFFFFF",
		Content:  "这里是标题文字",
		FontPath: "./assets/fzht.ttf",
	}

	qrCodeDraw := &QRCodeDraw{
		X:       30,
		Y:       860,
		Size:    250,
		Content: "https://www.minzhan.com",
	}

	// 创建流水线并添加组件
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		AddDrawer(avatarDraw).
		AddDrawer(textDraw).
		AddDrawer(qrCodeDraw).
		SetOutput(getOutputPath())

	// 执行绘制流程
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}

	fmt.Println("Pipeline execution success")
}

// TestPosterBuilderExample 展示如何使用PosterBuilder
func TestPosterBuilderExample(t *testing.T) {
	// 使用PosterBuilder创建海报
	builder := NewPosterBuilder(750, 1334)

	// 添加绘制组件
	builder.AddDrawer(
		&ImageDraw{
			X:    0,
			Y:    0,
			Path: "./assets/template.png",
		},
		&ImageDraw{
			X:     30,
			Y:     50,
			Path:  "./assets/avatar.jpg",
			Round: true,
		},
		&TextDraw{
			X:        180,
			Y:        105,
			Size:     26,
			Color:    "#FFFFFF",
			Content:  "这里是大标题1",
			FontPath: "./assets/fzht.ttf",
		},
		&TextDraw{
			X:        180,
			Y:        150,
			Size:     20,
			Color:    "#FFFFFF",
			Content:  "这里是小标题2",
			FontPath: "./assets/fzht.ttf",
		},
		&QRCodeDraw{
			X:       30,
			Y:       860,
			Size:    250,
			Content: "https://www.minzhan.com",
		},
	)
	// 构建并保存海报
	image, err := builder.Build()
	if err != nil {
		t.Errorf("Build poster failed: %v", err)
		return
	}
	// 保存图片
	outputPath := getOutputPath()
	if err = imagex.SavePNG(image, outputPath); err != nil {
		t.Errorf("Save image failed: %v", err)
		return
	}
	fmt.Println("Poster builder success", outputPath)
}

// TestMultiLineTextExample 展示多行文本绘制
func TestMultiLineTextExample(t *testing.T) {
	templatePath := "./assets/white.png"
	width, height, err := imagex.GetSize(templatePath)
	if err != nil {
		t.Skipf("Skip test due to missing template: %v", err)
		return
	}
	// 创建上下文
	ctx := NewContext(width, height)
	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: templatePath,
	}
	// 绘制文字
	textDraw1 := &TextDraw{
		X:        132,
		Y:        190,
		Size:     28,
		Color:    "#A5A6A8",
		Content:  "Minzhan All Rights Reserved.",
		FontPath: "./assets/fzht.ttf",
	}
	textDraw2 := &TextDraw{
		X:        132,
		Y:        236,
		Size:     30,
		Color:    "#A5A6A8",
		Content:  "民站科技（深圳）有限公司",
		FontPath: "./assets/fzht.ttf",
	}
	// 单行文本居左对齐
	textDraw3 := &MultiLineTextDraw{
		X:        20,
		Y:        300,
		AX:       FlexStart,
		AY:       FlexCenter,
		Size:     30,
		Color:    "#999999",
		Content:  "居左对齐",
		FontPath: "./assets/fzht.ttf",
	}
	// 单行文本居中对齐
	textDraw4 := &MultiLineTextDraw{
		X:           float64(width / 2),
		Y:           330,
		AX:          FlexCenter,
		AY:          FlexCenter,
		Size:        30,
		Color:       "#999999",
		Content:     "居中对齐居中对齐居中对齐",
		FontPath:    "./assets/fzht.ttf",
		CorrectionY: -5,
	}
	// 单行文本居右对齐
	textDraw5 := &MultiLineTextDraw{
		X:           float64(width - 20),
		Y:           360,
		AX:          FlexEnd,
		AY:          FlexCenter,
		Size:        30,
		Color:       "#999999",
		Content:     "居右对齐",
		FontPath:    "./assets/fzht.ttf",
		CorrectionY: -5,
	}
	// 多行文本居左对齐
	textDraw6 := &MultiLineTextDraw{
		X:           20,
		Y:           400,
		AX:          FlexStart,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居左对齐\n第二行文本\n第三行文本",
		FontPath:    "./assets/fzht.ttf",
		Align:       AlignLeft,
		LineSpacing: 1,
		CorrectionY: -5,
	}
	// 多行文本居中对齐
	textDraw7 := &MultiLineTextDraw{
		X:           float64(width / 2),
		Y:           550,
		AX:          FlexCenter,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居中对齐\n第二行文本\n第三行文本",
		FontPath:    "./assets/fzht.ttf",
		Align:       AlignCenter,
		LineSpacing: 2,
		CorrectionY: -5,
	}
	// 多行文本居右对齐
	textDraw8 := &MultiLineTextDraw{
		X:           float64(width) - 20,
		Y:           700,
		AX:          FlexEnd,
		Size:        30,
		Color:       "#FF0099",
		Content:     "多行文本居右对齐\n第二行文本\n第三行文本",
		FontPath:    "./assets/fzht.ttf",
		Align:       AlignRight,
		LineSpacing: 2,
		CorrectionY: -5,
	}

	// 创建流水线
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
		SetOutput(getOutputPath())
	// 执行绘制
	if err = pipeline.Execute(ctx); err != nil {
		t.Errorf("Multi-line text test failed: %v", err)
		return
	}
	fmt.Println("Multi-line text test success")
}

// TestJSONConfigExample 展示如何从JSON配置创建海报
func TestJSONConfigExample(t *testing.T) {
	jsonConfig := `[
		{
			"type": "image",
			"data": {
				"x": 0,
				"y": 0,
				"path": "./assets/white.png"
			}
		},
		{
			"type": "text",
			"data": {
				"x": 100,
				"y": 100,
				"size": 30,
				"color": "#000000",
				"content": "从JSON配置创建的文本",
				"fontPath": "./assets/fzht.ttf"
			}
		}
	]`
	builder := NewPosterBuilder(750, 1100)
	builder, err := builder.FromJSONConfig(jsonConfig)
	if err != nil {
		t.Errorf("Failed to load from JSON config: %v", err)
		return
	}
	_, err = builder.Build()
	if err != nil {
		t.Errorf("Failed to build from JSON config: %v", err)
		return
	}
	fmt.Println("JSON config test success")
}

// TestContent 展示如何绘制包含大量文本内容的海报
func TestContent(t *testing.T) {
	// 创建绘制上下文
	ctx := NewContext(750, 1334)
	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}
	//avatarDraw := &ImageDraw{
	//	X:     30,
	//	Y:     50,
	//	Path:  "https://minzhan.net/uploads/image/avatar.png",
	//	Round: true,
	//}
	textDraw := &TextDraw{
		X:        180,
		Y:        150,
		Size:     20,
		Color:    "#FFFFFF",
		Content:  "APP开发需要注意的细节非常多，这里罗列一些，避免大家踩坑：\n1、找开发公司或者APP开发团队。要多渠道的找，找一些觉得靠谱的开发公司，多接触，创始人最好是技术出身，有技术基因的公司才能保证项目的开发质量。很多的做销售出身的公司，技术能力真心没有办法保证。\n2、合同签订。合同内容一定要细致，需要有比较详细的列表和功能描述，这样才能保证后期不会出现扯皮。因为软件开发需求经常会变，开发公司有时候也有愉懒的情况。\n3、需求沟通。沟通结果一定要落实到纸或者邮件、文档。最后要产生详尽的产品原型。原型是必须的，产品文档可根据实际情况来确定要不要，因为产品文档这个太需要时间，可能咱的费用及开发公司精力等方面限制，PRD文档不是必须的。\n4、产品研发。一定要提前沟通好技术架构，这样对项目开发内沟通，以及后续产品版本迭代都会有非常大的帮助，减少沟通成本，提高开发效率和质量。\n5、产品测试。这个环节非常重要，咱们需要在beta版本的时候参与进来。这样可以更早的了解熟悉软件的实现情况，为后续运营作好准备。\n6、产品验收。一定要把各个功能细节，一定都要过2~5遍。这样双方都放心一些。 ",
		FontPath: "./assets/fzht.ttf",
	}
	// 创建流水线并添加组件
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		//AddDrawer(avatarDraw).
		AddDrawer(textDraw).
		SetOutput(getOutputPath())
	// 执行绘制流程
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}
	fmt.Println("Content test success")
}

// TestCircle 展示如何绘制圆形图片
func TestCircle(t *testing.T) {
	srcImage, err := imagex.ReadImage("./assets/template.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	newImage := DrawCircle(srcImage)
	if err = imagex.SavePNG(newImage, getOutputPath()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Circle test success")
}

// TestPoster 展示如何绘制包含多个组件的海报
func TestPoster(t *testing.T) {
	// 创建绘制上下文
	ctx := NewContext(750, 1334)
	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}
	qrCodeDraw := &QRCodeDraw{
		X:       30,
		Y:       860,
		Size:    250,
		Content: "https://www.minzhan.com",
	}
	avatarDraw := &ImageDraw{
		X:     30,
		Y:     250,
		Path:  "./assets/avatar.jpg",
		Round: true,
	}
	textDraw1 := &TextDraw{
		X:        180,
		Y:        105,
		Size:     26,
		Color:    "#FFFFFF",
		Content:  "这里是大标题1",
		FontPath: "./assets/fzht.ttf",
	}
	textDraw2 := &TextDraw{
		X:        180,
		Y:        150,
		Size:     20,
		Color:    "#FFFFFF",
		Content:  "这里是小标题2",
		FontPath: "./assets/fzht.ttf",
	}
	// 创建流水线并添加组件
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		AddDrawer(avatarDraw).
		AddDrawer(textDraw1).
		AddDrawer(textDraw2).
		AddDrawer(qrCodeDraw).
		SetOutput(getOutputPath())
	// 执行绘制流程
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}
	fmt.Println("Poster test success")
}

// TestGoodsPoster 展示如何绘制商品海报
func TestGoodsPoster(t *testing.T) {
	// 创建绘制上下文
	ctx := NewContext(750, 1100)
	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/white.png",
	}
	avatarDraw1 := &ImageDraw{
		X:     30,
		Y:     30,
		Path:  "https://static.minzhan.com/uploads/s274600676091367425/thumb/202406/1331ff9d1377f549450280b7509786308ad3.webp",
		Width: 690,
		Round: false,
	}
	avatarDraw2 := &ImageDraw{
		X:     500,
		Y:     780,
		Path:  "https://res.wx.qq.com/wxdoc/dist/assets/img/skyline-demo.37eff20b.png",
		Width: 200,
		Round: true,
	}
	textDraw1 := &TextDraw{
		X:        480,
		Y:        1020,
		Size:     30,
		Color:    "#000000",
		Content:  "长按识别小程序码",
		FontPath: "./assets/fzht.ttf",
	}
	textDraw2 := &TextDraw{
		X:        30,
		Y:        800,
		Size:     36,
		Color:    "#000000",
		Content:  "马面裙 白色上衣红色织金妆花双工艺面料马面裙",
		FontPath: "./assets/fzht.ttf",
	}
	textDraw3 := &TextDraw{
		X:        30,
		Y:        860,
		Size:     40,
		Color:    "#FF0000",
		Content:  "100元",
		FontPath: "./assets/fzht.ttf",
	}
	// 创建流水线并添加组件
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		AddDrawer(avatarDraw1).
		AddDrawer(avatarDraw2).
		AddDrawer(textDraw1).
		AddDrawer(textDraw2).
		AddDrawer(textDraw3).
		SetOutput(getOutputPath())
	// 执行绘制流程
	if err := pipeline.Execute(ctx); err != nil {
		t.Errorf("Pipeline execution failed: %v", err)
		return
	}
	fmt.Println("Goods poster test success")
}

func getOutputPath() string {
	outputPath := "d:/output/" + time.Now().Format("20060102150405") + ".png"
	_ = filex.MkdirAll(outputPath)
	return outputPath
}

// TestCustomPoster 展示如何绘制自定义海报
func TestCustomPoster(t *testing.T) {
	templatePath := "C:\\Users\\Administrator\\Desktop\\template.png"
	fontPath := "C:\\Users\\Administrator\\Desktop\\syht.ttf"
	width, height, err := imagex.GetSize(templatePath)
	if err != nil {
		t.Skipf("Skip test due to missing template: %v", err)
		return
	}
	// 创建绘制上下文
	ctx := NewContext(width, height)
	// 创建绘制组件
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: templatePath,
	}
	// 多行文本居左对齐
	textDraw1 := &MultiLineTextDraw{
		X:           130,
		Y:           160,
		AX:          FlexStart,
		Size:        34,
		Color:       "#A0A0A0",
		Content:     "Shell (China) Limited\n壳牌  (中国)  有限公司",
		FontPath:    fontPath,
		LineSpacing: 1.2,
		Align:       AlignLeft,
		CorrectionY: -5,
	}
	// 多行文本居左对齐
	textDraw2 := &MultiLineTextDraw{
		X:           130,
		Y:           479,
		AX:          FlexStart,
		Size:        29,
		Color:       "#888888",
		Content:     "产品编码\n产品型号\n批号\n生产日期",
		FontPath:    fontPath,
		LineSpacing: 2.7,
		Align:       AlignLeft,
		CorrectionY: -5,
	}
	// 多行文本居右对齐
	textDraw3 := &MultiLineTextDraw{
		X:           float64(width) - 130,
		Y:           479,
		AX:          FlexEnd,
		Size:        29,
		Color:       "#AAAAAA",
		Content:     "550045926\n壳牌液压油 TELLUS S1M68#200 升/桶\n12476898C60315JUL25\n2025-7-15-00:00:00",
		FontPath:    fontPath,
		LineSpacing: 2.7,
		Align:       AlignRight,
		CorrectionY: -5,
	}
	// 多行文本居左对齐
	textDraw4 := &MultiLineTextDraw{
		X:           130,
		Y:           884,
		AX:          FlexStart,
		Size:        34,
		Color:       "#212121",
		Content:     "This is to certify that the product(s) bearing the above\nbatch no(s) meets Shell's manufacturing/supply spec\nifications.\n特此证明以上批号的产品符合壳牌生产及供应规范。",
		FontPath:    fontPath,
		LineSpacing: 1.4,
		Align:       AlignLeft,
		CorrectionY: -5,
	}
	// 多行文本居左对齐
	textDraw5 := &MultiLineTextDraw{
		X:           130,
		Y:           1120,
		AX:          FlexStart,
		Size:        28,
		Color:       "#A0A0A0",
		Content:     "Shell (China) Limited\n壳牌  (中国)  有限公司",
		FontPath:    fontPath,
		LineSpacing: 1.2,
		Align:       AlignLeft,
		CorrectionY: -5,
	}
	// 创建流水线并添加组件
	pipeline := NewPipeline().
		AddDrawer(backgroundDraw).
		AddDrawer(textDraw1).
		AddDrawer(textDraw2).
		AddDrawer(textDraw3).
		AddDrawer(textDraw4).
		AddDrawer(textDraw5).
		SetOutput(getOutputPath())
	// 执行绘制流程
	if err = pipeline.Execute(ctx); err != nil {
		t.Errorf("Custom poster execution failed: %v", err)
		return
	}
	fmt.Println("Custom poster execution success")
}
