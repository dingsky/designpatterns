package builderPatterns
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director{
	return &Director{builder:builder}
}

func (d Director) Create() Builder  {
	d.builder.SetCharacter(d.builder.SetArms())
	d.builder.SetModel()
	return d.builder
}
