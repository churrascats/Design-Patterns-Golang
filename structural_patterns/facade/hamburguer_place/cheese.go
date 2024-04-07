package hamburguerplace

type Cheese struct {
	ingredient string
}

func NewCheese() *Cheese {
	return &Cheese{"Cheese"}
}

func (l *Cheese) AddCheese() string {
	return " " + l.ingredient
}
