package main

import (
	"fmt"
	"strategy/strategy"
)

func main() {
	context := strategy.NewContext()
	array := []int{4, 7, 1, 9, 7, 5, 2, 5, 4, 2}

	strategyA := strategy.SortedArrayStrategy{}
	strategyB := strategy.SortedReverseArrayStrategy{}
	strategyC := strategy.SortedDoubledArrayStrategy{}

	context.SetStrategy(&strategyA)
	fmt.Println("Showing arrays items for strategy A")
	context.ShowArrayItems(array)

	context.SetStrategy(&strategyB)
	fmt.Println("Showing arrays items for strategy B")
	context.ShowArrayItems(array)

	context.SetStrategy(&strategyC)
	fmt.Println("Showing arrays items for strategy C")
	context.ShowArrayItems(array)
}
