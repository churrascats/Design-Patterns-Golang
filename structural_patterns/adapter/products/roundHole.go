package products

type RoundHole struct {
	Radius float64
}

func (r *RoundHole) Fits(roundPeg IRoundPeg) bool {
	return r.Radius >= roundPeg.GetRadius()
}
