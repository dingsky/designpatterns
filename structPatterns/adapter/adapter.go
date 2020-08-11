package adapter

import (
	"fmt"
)

type V110ComputerInterface interface {
	Run110VComputer(voltage int64)
}

type V110Computer struct {}

func (m *V110Computer)Run110VComputer(voltage int64){
	if voltage != 110 {
		fmt.Printf("Can only run on 110V voltage!\n")
		return
	}
	fmt.Printf("Running 110V Computer!\n")
}

type V220ComputerInterface interface {
	Run220VComputer(voltage int64)
}

type V220Computer struct {}

func (p *V220Computer)Run220VComputer(voltage int64) {
	if voltage != 220 {
		fmt.Printf("Can only run on 220V voltage!\n")
		return
	}
	fmt.Printf("Running 220V Computer!\n")}

type Adapter struct {
	V110ComputerInterface
}

func (a *Adapter)Run220VComputer(voltage int64) {
	a.Run110VComputer(voltage/2)
}

func NewAdapter(V110Computer V110ComputerInterface) V220ComputerInterface {
	return &Adapter{
		V110ComputerInterface:V110Computer,
	}
}
