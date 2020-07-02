package simpleFactory

const (
	Cockoo = "cockoo"
	Seagull = "seagull"
	Eagle = "eagle"
)

// bird 鸟类
type bird interface {
	Fly()
	Sing()
	Eat()
}

func NewBird(birdName string) bird {
	switch birdName {
	case Cockoo:
		return new(cuckoo)
	case Seagull:
		return new(seagull)
	default:
		return new(eagle)
	}
}

