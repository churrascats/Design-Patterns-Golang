package main

import (
	"fmt"
	"lightweight/game"
)

func main() {
	newGame := game.NewGame()

	newGame.AddTerrorist(game.TerroristDressType)
	newGame.AddTerrorist(game.TerroristDressType)
	newGame.AddTerrorist(game.TerroristDressType)
	newGame.AddTerrorist(game.TerroristDressType)

	newGame.AddCounterTerrorist(game.CounterTerrroristDressType)
	newGame.AddCounterTerrorist(game.CounterTerrroristDressType)
	newGame.AddCounterTerrorist(game.CounterTerrroristDressType)

	dressFactoryInstance := game.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
