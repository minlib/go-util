package poster

import (
	"image"
)

// Context Context
type Context struct {
	Canvas *image.RGBA
}

type IDraw interface {
	Do(c *Context) error          // 自身的业务
	SetNext(draws ...IDraw) IDraw // 设置下一步对象
	Run(c *Context) error         // 执行
}

type NextStep struct {
	nextDraw IDraw
}

func (n *NextStep) SetNext(draws ...IDraw) IDraw {
	if len(draws) == 0 {
		return nil
	}
	n.nextDraw = draws[0]
	if len(draws) == 0 {
		return draws[0]
	}
	for i := 0; i < len(draws)-1; i++ {
		draws[i].SetNext(draws[i+1])
	}
	return draws[len(draws)-1]
}

func (n *NextStep) Run(c *Context) error {
	if n.nextDraw != nil {
		if err := n.nextDraw.Do(c); err != nil {
			return err
		}
		return n.nextDraw.Run(c)
	}
	return nil
}
