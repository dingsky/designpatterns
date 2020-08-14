package filter

const (
	Man = "男"
	Woman = "女"

	PrimarySchool = "小学"
	JuniorMiddleSchool = "初中"
	HighSchool = "高中"
)

type Student struct {
	Name string
	Sex string
	Grade string
}

func NewStudent(name, sex, grade string) Student {
	return Student{Name:name, Sex:sex, Grade: grade}
}