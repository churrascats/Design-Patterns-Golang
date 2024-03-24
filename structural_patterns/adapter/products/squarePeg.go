package products

type SquarePeg struct {
	Size float64
}

func (s *SquarePeg) GetSquareSize() float64 {
	return s.Size
}
