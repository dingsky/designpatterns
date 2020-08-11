package builderPatterns

import "fmt"

type Arm struct {
	ArmName   string  // 武器名称
	ArmWeight float64 // 武器重量
	ArmCount  int64   // 武器数量
}

type ArmList struct {
	Arms []*Arm
}

func (a *ArmList) AddArm(arm *Arm) {
	a.Arms = append(a.Arms, arm)
}

func (a *ArmList) ShowArms() string {
	if len(a.Arms) == 0 {
		return "没有装备武器"
	}

	result := fmt.Sprintf("武器列表:\n")
	for _, arm := range a.Arms {
		armShow := fmt.Sprintf("武器名称:%s 武器重量:%f 武器数量:%s\n", arm.ArmName, arm.ArmWeight, arm.ArmCount)
		result += armShow
	}
	return result
}

func (a *ArmList) GetArmsTotalWeight() float64 {
	var result float64
	for _, arm := range a.Arms {
		result += arm.ArmWeight
	}
	return result
}
