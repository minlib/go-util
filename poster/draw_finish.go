package poster

import (
	"github.com/minlib/go-util/imagex"
)

// FinishDraw finish and output image
type FinishDraw struct {
	Output string
	NextStep
}

func (d *FinishDraw) Do(c *Context) error {
	if err := imagex.SavePNG(c.Canvas, d.Output); err != nil {
		return err
	}
	return nil
}
