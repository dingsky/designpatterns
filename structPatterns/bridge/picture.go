package bridge

import "fmt"

type Picture struct {
	color Color
	shape shape
}

func NewPicture(color Color, shape shape) *Picture {
	return &Picture{color: color, shape: shape}
}

func (p *Picture)Show() {
	fmt.Printf("我是：%s%s\n", p.color.Color(), p.shape.Shape())
}
