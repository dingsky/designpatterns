package simpleFactory

import "fmt"

// cuckoo 布谷鸟
type cuckoo struct {
}

func (*cuckoo) Fly() {
	fmt.Printf("Cuckoo flying!\n")
}

func (*cuckoo) Sing() {
	fmt.Printf("Cuckoo singing!\n")
}

func (*cuckoo) Eat() {
	fmt.Printf("I like eating worm!\n")
}
