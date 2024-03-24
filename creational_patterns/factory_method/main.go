package main

import (
	"factory_method/factory"
	"fmt"
)

func main() {
	polygonFactory := factory.PolygonFactory{}

	polygon := polygonFactory.GetPolygon(factory.CIRCLE)
	if polygon == nil {
		return
	}

	numberOfSides := polygon.GetNumberOfSides()
	fmt.Println(numberOfSides)
}
