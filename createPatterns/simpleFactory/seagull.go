package simpleFactory

import "fmt"

// seagull 布谷鸟
type seagull struct {
}

func(*seagull) Fly() {
	fmt.Printf("Seagull flying!\n")
}

func(*seagull) Sing() {
	fmt.Printf("Seagull singing!\n")
}

func(*seagull) Eat() {
	fmt.Printf("I like eating fish!\n")
}