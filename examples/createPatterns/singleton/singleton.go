package main

import (
	"designpatterns/createPatterns/singleton"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(singleton.GetInstance().SingletonData())
		time.Sleep(time.Second)
	}
}
