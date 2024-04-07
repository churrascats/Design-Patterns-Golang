package hamburguerplace

type Letucce struct {
	ingredient string
}

func NewLetucce() *Letucce {
	return &Letucce{"Letucce"}
}

func (l *Letucce) AddLetucce() string {
	return " " + l.ingredient
}
