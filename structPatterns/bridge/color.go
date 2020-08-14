package bridge

// 颜色
type Color interface{
	Color()string
}

// 红色
type Red struct {}

func(c *Red)Color() string{
	return "红色"
}


// 黄色
type Yellow struct {}

func(s *Yellow)Color() string{
	return "黄色"
}

// 蓝色
type Blue struct {}

func(r *Blue)Color() string{
	return "蓝色"
}
