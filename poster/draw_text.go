package poster

type TextDraw struct {
	FontPath string  //字体路径
	Content  string  //内容
	Color    string  //颜色
	Size     float64 //大小
	X        int     //横坐标
	Y        int     //纵坐标
	NextStep
}

func (d *TextDraw) Do(c *Context) error {
	if d.Size == 0 {
		d.Size = 24
	}
	text := &Text{
		Canvas:   c.Canvas,
		FontPath: d.FontPath,
		Content:  d.Content,
		Color:    d.Color,
		Size:     d.Size,
		X:        d.X,
		Y:        d.Y,
	}
	err := DrawText(text)
	if err != nil {
		return err
	}
	return nil
}
