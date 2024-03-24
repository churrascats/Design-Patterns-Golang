package visitor

type Rectangle struct {
	L int
	B int
}

func (r *Rectangle) GetType() string {
	return "rectangle"
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForrectangle(r)
}
