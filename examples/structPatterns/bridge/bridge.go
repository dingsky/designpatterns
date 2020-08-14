package main

import "designpatterns/structPatterns/bridge"

func main() {
	// 红色正方形
	redSquare := bridge.NewPicture(&bridge.Red{}, &bridge.Square{})
	redSquare.Show()

	// 蓝色圆形
	blueCircular := bridge.NewPicture(&bridge.Blue{}, &bridge.Circular{})
	blueCircular.Show()

	// 黄色长方形
	yellowRectangle := bridge.NewPicture(&bridge.Yellow{}, &bridge.Rectangle{})
	yellowRectangle.Show()

}
