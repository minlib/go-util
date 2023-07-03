package poster

import (
	"image"
)

// Context Context
type Context struct {
	Canvas *image.RGBA
}

type IDraw interface {
	Do(c *Context) error   // 自身的业务
	SetNext(h IDraw) IDraw // 设置下一步对象
	Run(c *Context) error  // 执行
}

type NextStep struct {
	nextDraw IDraw
}

func (n *NextStep) SetNext(h IDraw) IDraw {
	n.nextDraw = h
	return h
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
