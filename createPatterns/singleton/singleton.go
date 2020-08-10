package singleton

type singleton struct {
	data string
}

var singletonInstance = new(singleton)

func init() {
	singletonInstance.data = "Hello World!"
}

func GetInstance() *singleton {
	return singletonInstance
}

func (s *singleton) SingletonData() string {
	return s.data
}
