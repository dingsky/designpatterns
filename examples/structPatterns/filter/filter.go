package main

import (
	"designpatterns/structPatterns/filter"
	"fmt"
)

func main() {
	var studentList []filter.Student
	studentList = append(studentList, filter.NewStudent("陈子昆", filter.Man, filter.PrimarySchool))
	studentList = append(studentList, filter.NewStudent("闫肃", filter.Man, filter.JuniorMiddleSchool))
	studentList = append(studentList, filter.NewStudent("陈寿", filter.Man, filter.HighSchool))

	studentList = append(studentList, filter.NewStudent("嫣红", filter.Woman, filter.PrimarySchool))
	studentList = append(studentList, filter.NewStudent("唐嫣", filter.Woman, filter.JuniorMiddleSchool))
	studentList = append(studentList, filter.NewStudent("鉴冰", filter.Woman, filter.HighSchool))

	filterPrimarySchool := new(filter.FilterPrimarySchool)
	fmt.Printf("小学生清单: %v\n", filterPrimarySchool.Filter(studentList))

	filterJuniorMiddleSchool := new(filter.FilterJuniorMiddleSchool)
	fmt.Printf("初中生清单: %v\n", filterJuniorMiddleSchool.Filter(studentList))

	filterJHighSchool := new(filter.FilterJHighSchool)
	fmt.Printf("高中生清单: %v\n", filterJHighSchool.Filter(studentList))

	filterMan := new(filter.FilterMan)
	fmt.Printf("男孩清单: %v\n", filterMan.Filter(studentList))

	filterWoman := new(filter.FilterWoman)
	fmt.Printf("女孩清单: %v\n", filterWoman.Filter(studentList))


	fmt.Printf("高中女生清单: %v\n", filterWoman.Filter(filterJHighSchool.Filter(studentList)))
	fmt.Printf("初中男生清单: %v\n", filterMan.Filter(filterJuniorMiddleSchool.Filter(studentList)))


}
