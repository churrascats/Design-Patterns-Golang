package builder

import "builder/pizza"

type PizzaBuilder struct {
	size    int
	cheese  bool
	sauce   bool
	sausage bool
	bacon   bool
}

func (pb *PizzaBuilder) SetSize(size int) {
	pb.size = size
}

func (pb *PizzaBuilder) SetCheese(hasCheese bool) {
	pb.cheese = hasCheese
}

func (pb *PizzaBuilder) SetSauce(hasSauce bool) {
	pb.sauce = hasSauce
}

func (pb *PizzaBuilder) SetSausage(hasSausage bool) {
	pb.sausage = hasSausage
}

func (pb *PizzaBuilder) SetBacon(hasBacon bool) {
	pb.bacon = hasBacon
}

func (pb *PizzaBuilder) Build() pizza.IPizza {
	return pizza.NewPizza(
		pb.size,
		pb.cheese,
		pb.sauce,
		pb.sausage,
		pb.bacon,
	)
}
