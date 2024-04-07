package hamburguerplace

type Tomato struct {
	ingredient string
}

func NewTomato() *Tomato {
	return &Tomato{"Tomato"}
}

func (l *Tomato) AddTomato() string {
	return " " + l.ingredient
}
