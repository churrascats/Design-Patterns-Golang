package hamburguerplace

type Patty struct {
	ingredient string
}

func NewPatty() *Patty {
	return &Patty{"Patty"}
}

func (l *Patty) AddPatty() string {
	return " " + l.ingredient
}
