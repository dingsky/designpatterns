package main

import "designpatterns/createPatterns/simpleFactory"

func main() {
	cockoo := simpleFactory.NewBird(simpleFactory.Cockoo)
	cockoo.Fly()
	cockoo.Sing()
	cockoo.Eat()

	seagull := simpleFactory.NewBird(simpleFactory.Seagull)
	seagull.Fly()
	seagull.Sing()
	seagull.Eat()

	eagle := simpleFactory.NewBird(simpleFactory.Eagle)
	eagle.Fly()
	eagle.Sing()
	eagle.Eat()
}
