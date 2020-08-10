package simpleFactory

import "fmt"

// eagle 布谷鸟
type eagle struct {
}

func (*eagle) Fly() {
	fmt.Printf("Eagle flying!\n")
}

func (*eagle) Sing() {
	fmt.Printf("Eagle singing!\n")
}

func (*eagle) Eat() {
	fmt.Printf("I like eating meat!\n")
}
