package poster

// FinishDraw finish and output image
type FinishDraw struct {
	Output string
	NextStep
}

func (d *FinishDraw) Do(c *Context) error {
	if err := SavePNG(c.Canvas, d.Output); err != nil {
		return err
	}
	return nil
}
