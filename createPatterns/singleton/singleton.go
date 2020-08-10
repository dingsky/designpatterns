package singleton

type singleton struct {
	data string
}

var Singleton = new(singleton)

func init() {
	Singleton.data = "Hello World!"
}

func (s *singleton)SingletonData() string {
	return s.data
}