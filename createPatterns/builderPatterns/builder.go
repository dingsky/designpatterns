package builderPatterns

type Builder interface {
	SetCharacter(weightBearing float64)
	SetArms() float64
	SetModel()
	Build() interface{}
}

type SniperBuilder struct {
	Character Character
	Arms ArmList
	Model Model
}

func (s *SniperBuilder)SetCharacter(weightBearing float64) {
	s.Character.SetName("张三")
	s.Character.SetAge(30)
	s.Character.SetHeight(1.8)
	s.Character.SetSex(Man)
	s.Character.SetWeightBearing(weightBearing)
	s.Character.SetSpeed()
}

func (s *SniperBuilder)SetArms() float64 {
	arm := new(Arm)
	arm.ArmWeight = 30
	arm.ArmName = SniperRifle
	arm.ArmCount = 1
	s.Arms.AddArm(arm)

	arm.ArmWeight = 0.5
	arm.ArmName = AntitankGrenade
	arm.ArmCount = 5
	s.Arms.AddArm(arm)

	arm.ArmWeight = 5
	arm.ArmName = Pistol
	arm.ArmCount = 1
	s.Arms.AddArm(arm)
	return s.Arms.GetArmsTotalWeight()
}

func (s *SniperBuilder)SetModel() {
	s.Model.SetName(CharacterModel1)
}

func (s *SniperBuilder)Build() interface{} {
	return s
}

type SMGInfantryBuilder struct {
	Character Character
	Arms ArmList
	Model Model
}

func (s *SMGInfantryBuilder)SetCharacter(weightBearing float64) {
	s.Character.SetName("李四")
	s.Character.SetAge(28)
	s.Character.SetHeight(1.75)
	s.Character.SetSex(Man)
	s.Character.SetWeightBearing(weightBearing)
	s.Character.SetSpeed()
}

func (s *SMGInfantryBuilder)SetArms() float64 {
	arm := new(Arm)
	arm.ArmWeight = 20
	arm.ArmName = SubmachineGun
	arm.ArmCount = 1
	s.Arms.AddArm(arm)

	arm.ArmWeight = 0.5
	arm.ArmName = AntitankGrenade
	arm.ArmCount = 10
	s.Arms.AddArm(arm)

	arm.ArmWeight = 5
	arm.ArmName = Pistol
	arm.ArmCount = 1
	s.Arms.AddArm(arm)
	return s.Arms.GetArmsTotalWeight()

}

func (s *SMGInfantryBuilder)SetModel() {
	s.Model.SetName(CharacterModel3)
}

func (s *SMGInfantryBuilder)Build() interface{} {
	return s
}

type HostageBuilder struct {
	Character Character
	Arms ArmList
	Model Model
}

func (s *HostageBuilder)SetCharacter(weightBearing float64) {
	s.Character.SetName("李娜")
	s.Character.SetAge(22)
	s.Character.SetHeight(1.65)
	s.Character.SetSex(Woman)
	s.Character.SetWeightBearing(weightBearing)
	s.Character.SetSpeed()
}

func (s *HostageBuilder)SetArms() float64 {
	return 0
}

func (s *HostageBuilder)SetModel() {
	s.Model.SetName(CharacterModel2)
}

func (s *HostageBuilder)Build() interface{} {
	return s
}
