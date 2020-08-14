package filter

type filter interface {
	Filter([]Student) []Student
}

type FilterMan struct {}

func (f *FilterMan)Filter(studentList []Student) []Student {
	studentResult := make([]Student, 0, len(studentList))
	for _, student := range studentList {
		if student.Sex == Man {
			studentResult = append(studentResult, student)
		}
	}
	return studentResult
}

type FilterWoman struct {}

func (f *FilterWoman)Filter(studentList []Student) []Student {
	studentResult := make([]Student, 0, len(studentList))
	for _, student := range studentList {
		if student.Sex == Woman {
			studentResult = append(studentResult, student)
		}
	}
	return studentResult
}

type FilterPrimarySchool struct {}

func (f *FilterPrimarySchool)Filter(studentList []Student) []Student {
	studentResult := make([]Student, 0, len(studentList))
	for _, student := range studentList {
		if student.Grade == PrimarySchool {
			studentResult = append(studentResult, student)
		}
	}
	return studentResult
}

type FilterJuniorMiddleSchool struct {}

func (f *FilterJuniorMiddleSchool)Filter(studentList []Student) []Student {
	studentResult := make([]Student, 0, len(studentList))
	for _, student := range studentList {
		if student.Grade == JuniorMiddleSchool {
			studentResult = append(studentResult, student)
		}
	}
	return studentResult
}

type FilterJHighSchool struct {}

func (f *FilterJHighSchool)Filter(studentList []Student) []Student {
	studentResult := make([]Student, 0, len(studentList))
	for _, student := range studentList {
		if student.Grade == HighSchool {
			studentResult = append(studentResult, student)
		}
	}
	return studentResult
}

