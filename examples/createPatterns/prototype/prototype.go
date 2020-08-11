package main

import (
	"designpatterns/createPatterns/prototype"
)

func main() {
	student := new(prototype.Student)
	student.Name = "sky"
	student.Sex = "man"
	student.Age = 18
	student.Curriculums = []string{"C语言", "JAVA", "编译原理", "数据库原理", "离散数学", "概率论"}

	student1 := student.Clone()
	student1.Name = "spring"
	student1.Age = 20

	student2 := student.Clone()
	student2.Name = "lucy"
	student2.Sex = "woman"
	student2.Age = 21

	student.Show()
	student1.Show()
	student2.Show()
}
