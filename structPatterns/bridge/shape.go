package bridge

// 形状
type shape interface{
	Shape()string
}

// 圆形
type Circular struct {}

func(c *Circular)Shape() string{
	return "圆形"
}


// 正方形
type Square struct {}

func(s *Square)Shape() string{
	return "正方形"
}

// 长方形
type Rectangle struct {}

func(r *Rectangle)Shape() string{
	return "长方形"
}
