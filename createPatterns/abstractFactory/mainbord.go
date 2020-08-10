package abstractFactory

import "fmt"

type mainbord interface {
	SetupMainBord()
}

func NewMainBord(mainbordType string) mainbord {
	switch mainbordType {
	case Dell:
		return new(dellMainBord)
	default:
		return new(lenovoMainBord)
	}
}

type dellMainBord struct{}

func (d *dellMainBord) SetupMainBord() {
	fmt.Printf("I am seting dell main bord\n")
}

type lenovoMainBord struct{}

func (d *lenovoMainBord) SetupMainBord() {
	fmt.Printf("I am seting lenovo main bord\n")
}
