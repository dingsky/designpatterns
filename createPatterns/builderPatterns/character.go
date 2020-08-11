package builderPatterns

const (
	CharacterModel1 = "狙击手模型"
	CharacterModel2 = "人质模型"
	CharacterModel3 = "步枪兵模型"
)

const (
	Pistol          = "手枪"
	SniperRifle     = "狙击步枪"
	SubmachineGun   = "冲锋枪"
	AntitankGrenade = "手雷"
)

const (
	Man   = "男性"
	Woman = "女性"
)

// 角色
type Character struct {
	name          string  // 姓名
	Height        float64 // 身高
	Age           int64   // 年龄
	Sex           string  // 性别
	Speed         float64 // 跑步速度
	WeightBearing float64 // 负重
}

func (c *Character) SetName(name string) {
	c.name = name
}

func (c *Character) GetName() string {
	return c.name
}

func (c *Character) SetHeight(height float64) {
	c.Height = height
}

func (c *Character) GetHeight() float64 {
	return c.Height
}

func (c *Character) SetAge(age int64) {
	c.Age = age
}

func (c *Character) GetAge() int64 {
	return c.Age
}

func (c *Character) SetSex(sex string) {
	c.Sex = sex
}

func (c *Character) GetSex() string {
	return c.Sex
}

func (c *Character) SetSpeed() {
	c.Speed = 100 / (c.WeightBearing + c.Height + float64(c.Age))
}

func (c *Character) GetSpeed() float64 {
	return c.Speed
}

func (c *Character) SetWeightBearing(weightBearing float64) {
	c.WeightBearing = weightBearing
}

func (c *Character) GetWeightBearing() float64 {
	return c.WeightBearing
}
