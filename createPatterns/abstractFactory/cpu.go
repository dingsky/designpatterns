package abstractFactory

import "fmt"

type cpu interface {
	SetUpCpu()
}

func NewCpu(cpuType string) cpu {
	switch cpuType {
	case Dell:
		return new(dellCpu)
	default:
		return new(lenovoCpu)
	}
}

type dellCpu struct{}

func (d *dellCpu) SetUpCpu() {
	fmt.Printf("I am seting dell cpu\n")
}

type lenovoCpu struct{}

func (d *lenovoCpu) SetUpCpu() {
	fmt.Printf("I am seting lenovo cpu\n")
}
