package products

type IRoundPeg interface {
	GetRadius() float64
}

type RoundPeg struct {
	Radius float64
}

func (r *RoundPeg) GetRadius() float64 {
	return r.Radius
}
