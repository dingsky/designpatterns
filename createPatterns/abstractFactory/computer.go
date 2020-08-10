package abstractFactory

type computer interface {
	CreateMainBord() mainbord
	CreateCpu() cpu
}

func CreateComputer(computerType string) computer {
	switch computerType {
	case Dell:
		return new(dellComputer)
	default:
		return new(lenovoComputer)
	}
}

type lenovoComputer struct {
}

func (l *lenovoComputer) CreateMainBord() mainbord {
	return NewMainBord(Lenovo)
}

func (l *lenovoComputer) CreateCpu() cpu {
	return NewCpu(Lenovo)
}

type dellComputer struct {
}

//
func (l *dellComputer) CreateMainBord() mainbord {
	return NewMainBord(Dell)
}

func (l *dellComputer) CreateCpu() cpu {
	return NewCpu(Dell)
}
