package main

import (
	"designpatterns/structPatterns/adapter"
)

// 使用220V的电压来运行110V的电脑, 中间相当于做了个变压
// 使用220V的新接口, 兼容了110V的老接口


func main() {
	v110Computer := new(adapter.V110Computer)
	RunComputer(v110Computer)
	v220Computer := new(adapter.V220Computer)
	RunComputer(v220Computer)
}

func RunComputer(computer interface{}) {
	switch computer.(type) {
	case adapter.V110ComputerInterface:
		v220Computer := adapter.NewAdapter(computer.(adapter.V110ComputerInterface))
		v220Computer.Run220VComputer(220)
	default:
		computer.(adapter.V220ComputerInterface).Run220VComputer(220)
	}
}
