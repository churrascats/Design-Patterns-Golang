package factory

type Square struct{}

func (s *Square) GetNumberOfSides() int {
	return 4
}
