package poster

import (
	"fmt"
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
	"testing"
	"time"
)

func TestContent(t *testing.T) {
	start := &StartDraw{}
	ctx := &Context{
		Canvas: NewRGBA(0, 0, 750, 1334),
	}
	// 绘制背景图
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}
	// 绘制远程头像
	avatarDraw1 := &ImageDraw{
		X:     30,
		Y:     50,
		Path:  "https://minzhan.net/uploads/image/avatar.png",
		Round: true,
	}
	// 绘制文字
	textDraw1 := &TextDraw{
		X:        180,
		Y:        150,
		Size:     20,
		Color:    "#FFFFFF",
		Content:  "APP开发需要注意的细节非常多，这里罗列一些，避免大家踩坑：\n1、找开发公司或者APP开发团队。要多渠道的找，找一些觉得靠谱的开发公司，多接触，创始人最好是技术出身，有技术基因的公司才能保证项目的开发质量。很多的做销售出身的公司，技术能力真心没有办法保证。\n2、合同签订。合同内容一定要细致，需要有比较详细的列表和功能描述，这样才能保证后期不会出现扯皮。因为软件开发需求经常会变，开发公司有时候也有愉懒的情况。\n3、需求沟通。沟通结果一定要落实到纸或者邮件、文档。最后要产生详尽的产品原型。原型是必须的，产品文档可根据实际情况来确定要不要，因为产品文档这个太需要时间，可能咱的费用及开发公司精力等方面限制，PRD文档不是必须的。\n4、产品研发。一定要提前沟通好技术架构，这样对项目开发内沟通，以及后续产品版本迭代都会有非常大的帮助，减少沟通成本，提高开发效率和质量。\n5、产品测试。这个环节非常重要，咱们需要在beta版本的时候参与进来。这样可以更早的了解熟悉软件的实现情况，为后续运营作好准备。\n6、产品验收。一定要把各个功能细节，一定都要过2~5遍。这样双方都放心一些。 ",
		FontPath: "./assets/fzht.ttf",
	}
	// 完成绘制，并导出图片
	finishDraw := &FinishDraw{
		Output: getOutputPath(),
	}
	start.
		SetNext(backgroundDraw).
		SetNext(avatarDraw1).
		SetNext(textDraw1).
		SetNext(finishDraw)
	if err := start.Run(ctx); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

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
	fmt.Println("Success")
}

func TestPoster(t *testing.T) {
	start := &StartDraw{}
	ctx := &Context{
		Canvas: NewRGBA(0, 0, 750, 1334),
	}
	// 绘制背景图
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}
	// 绘制二维码
	qrCodeDraw := &QRCodeDraw{
		X:       30,
		Y:       860,
		Size:    250,
		Content: "https://www.minzhan.com",
	}
	// 绘制本地头像
	avatarDraw2 := &ImageDraw{
		X:     30,
		Y:     250,
		Path:  "./assets/avatar.jpg",
		Round: true,
	}
	//// 绘制远程头像
	//avatarDraw1 := &ImageDraw{
	//	X:     30,
	//	Y:     50,
	//	Path:  "https://minzhan.net/uploads/image/avatar.png",
	//	Round: true,
	//}
	// 绘制文字
	textDraw1 := &TextDraw{
		X:        180,
		Y:        105,
		Size:     26,
		Color:    "#FFFFFF",
		Content:  "这里是大标题1",
		FontPath: "./assets/fzht.ttf",
	}
	// 绘制文字
	textDraw2 := &TextDraw{
		X:        180,
		Y:        150,
		Size:     20,
		Color:    "#FFFFFF",
		Content:  "这里是小标题2",
		FontPath: "./assets/fzht.ttf",
	}
	// 完成绘制，并导出图片
	finishDraw := &FinishDraw{
		Output: getOutputPath(),
	}
	start.
		SetNext(backgroundDraw).
		//SetNext(avatarDraw1).
		SetNext(avatarDraw2).
		SetNext(textDraw1).
		SetNext(textDraw2).
		SetNext(qrCodeDraw).
		SetNext(finishDraw)
	if err := start.Run(ctx); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func getOutputPath() string {
	fmt.Println(time.Now().UnixMicro())
	outputPath := "d:/output/" + time.UnixMilli(time.Now().UnixMilli()).Format("20060102150405.000") + ".png"
	_ = filex.MkdirAll(outputPath)
	return outputPath
}
