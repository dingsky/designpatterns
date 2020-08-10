package main

import "designpatterns/createPatterns/abstractFactory"

// 虽然不需要知道每个工厂的具体实现, 但是还是要知道每个抽象工厂包含哪些工厂实例, 每个实例对应的调用方法.
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
