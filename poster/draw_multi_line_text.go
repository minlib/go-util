package poster

type MultiLineTextDraw struct {
	FontPath    string    // 字体路径
	Content     string    // 内容
	Color       string    // 颜色
	Size        float64   // 大小
	X           float64   // 横坐标
	Y           float64   // 纵坐标
	AX          float64   // 水平锚点比例（0左，0.5中，1右）
	AY          float64   // 垂直锚点比例（0上，0.5中，1下）
	LineSpacing float64   // 行间距
	Align       TextAlign // 对齐方式
	NextStep
}

func (d *MultiLineTextDraw) Do(c *Context) error {
	if d.Size == 0 {
		d.Size = 24
	}
	multiLineText := &MultiLineText{
		Content:     d.Content,
		FontPath:    d.FontPath,
		Color:       d.Color,
		Size:        d.Size,
		X:           d.X,
		Y:           d.Y,
		AX:          d.AX,
		AY:          d.AY,
		LineSpacing: d.LineSpacing,
		Align:       d.Align,
	}
	err := DrawMultiLineText(c.Canvas, multiLineText)
	if err != nil {
		return err
	}
	return nil
}
