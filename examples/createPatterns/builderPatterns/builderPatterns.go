package main

import (
	"designpatterns/createPatterns/builderPatterns"
	"fmt"
)

func main () {
	sniperBuilder := new(builderPatterns.SniperBuilder)
	directpor := builderPatterns.NewDirector(sniperBuilder)
	directpor.Create()
	fmt.Printf("%v\n", sniperBuilder)

	smgInfantryBuilder := new(builderPatterns.SMGInfantryBuilder)
	directpor = builderPatterns.NewDirector(smgInfantryBuilder)
	directpor.Create()
	fmt.Printf("%v\n", smgInfantryBuilder)

	hostageBuilder := new(builderPatterns.HostageBuilder)
	directpor = builderPatterns.NewDirector(hostageBuilder)
	directpor.Create()
	fmt.Printf("%v\n", smgInfantryBuilder)
}
