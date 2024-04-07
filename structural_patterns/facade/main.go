package main

import (
	"facade/facade"
	"fmt"
)

func main() {
	hamburguerFacade := facade.NewHamburguerFacade()

	hamburguer := hamburguerFacade.GetHamburguer()

	fmt.Println(hamburguer)
}
