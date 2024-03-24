package pizza

import (
	"fmt"
)

type IPizza interface {
	PrintRecipe()
}

type Pizza struct {
	size    int
	cheese  bool
	sauce   bool
	sausage bool
	bacon   bool
}

func NewPizza(
	size int,
	cheese,
	sauce,
	sausage,
	bacon bool) IPizza {
	return Pizza{
		size:    size,
		cheese:  cheese,
		sauce:   sauce,
		sausage: sausage,
		bacon:   bacon,
	}
}

func (p Pizza) PrintRecipe() {
	pizzaRecipe := fmt.Sprintf(
		"Size: %v, cheese: %t, sauce: %t, sausage: %t, bacon: %t",
		p.size, p.cheese, p.sauce, p.sausage, p.bacon,
	)

	fmt.Println(pizzaRecipe)
}
