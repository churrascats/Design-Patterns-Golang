package factory

type Triangle struct {
}

func (t *Triangle) GetNumberOfSides() int {
	return 3
}
