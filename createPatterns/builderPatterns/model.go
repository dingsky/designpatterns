package builderPatterns

type Model struct {
	ModelName string
}

func (m *Model) SetName(name string) {
	m.ModelName = name
}

func (m *Model) GetName() string {
	return m.ModelName
}
