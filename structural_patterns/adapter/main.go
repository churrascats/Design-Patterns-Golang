package main

import (
	"adapter/adapters"
	"adapter/products"
	"fmt"
)

func main() {
	roundHole := products.RoundHole{Radius: 4}
	roundPeg := &products.RoundPeg{Radius: 2}

	itFits := roundHole.Fits(roundPeg)
	fmt.Println(itFits)

	squarePeg := products.SquarePeg{Size: 2}
	squarePegAdapter := adapters.SquarePegAdapter{SquarePeg: squarePeg}

	itFits = roundHole.Fits(&squarePegAdapter)
	fmt.Println(itFits)

}
