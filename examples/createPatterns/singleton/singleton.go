package main

import (
	"designpatterns/createPatterns/singleton"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(singleton.GetHungryManSingleton().SingletonData())
		time.Sleep(time.Second)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(singleton.GetLazyManSingleton().SingletonData())
		time.Sleep(time.Second)
	}
}
