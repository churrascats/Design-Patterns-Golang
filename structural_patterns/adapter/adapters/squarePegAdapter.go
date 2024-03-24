package adapters

import (
	"adapter/products"
	"math"
)

type SquarePegAdapter struct {
	products.RoundPeg
	SquarePeg products.SquarePeg
}

func (s *SquarePegAdapter) GetRadius() float64 {
	squareSize := s.SquarePeg.GetSquareSize()

	return math.Pow(squareSize, 2)
}
