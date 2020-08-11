package prototype

import "fmt"

// 定义原型接口, 具备自我克隆能力的对象都可以是原型对象
type Prototyle interface {
	Clone() Prototyle
}

type Student struct {
	Name        string   // 姓名
	Age         int64    // 年龄
	Sex         string   // 性别
	Curriculums []string // 课程
}

func (s *Student) Clone() *Student {
	student := *s
	return &student
}

func (s *Student) Show() {
	fmt.Printf("姓名:%s 年龄:%d 性别:%s 课程:%v\n", s.Name, s.Age, s.Sex, s.Curriculums)
}
