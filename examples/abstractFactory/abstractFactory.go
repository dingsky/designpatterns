package main

import "designpatterns/createPatterns/abstractFactory"

func main() {
	// 创建抽象工厂
	computer := abstractFactory.CreateComputer(abstractFactory.Dell)

	// 创建Cpu
	cpu := computer.CreateCpu()

	// 安装Cpu
	cpu.SetUpCpu()

	// 创建主板
	mainbord := computer.CreateMainBord()

	// 安装主板
	mainbord.SetupMainBord()

	// 创建抽象工厂
	computer = abstractFactory.CreateComputer(abstractFactory.Lenovo)

	// 创建Cpu
	cpu = computer.CreateCpu()

	// 安装Cpu
	cpu.SetUpCpu()

	// 创建主板
	mainbord = computer.CreateMainBord()

	// 安装主板
	mainbord.SetupMainBord()
}
