package decorators

import "decorator/pizza"

type FrangoDecorator struct {
	Pizza pizza.Pizza
}

func (f *FrangoDecorator) GetPrice() int {
	return f.Pizza.GetPrice() + 20
}
