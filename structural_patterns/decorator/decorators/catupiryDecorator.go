package decorators

import "decorator/pizza"

type CatupiryDecorator struct {
	Pizza pizza.Pizza
}

func (c *CatupiryDecorator) GetPrice() int {
	return c.Pizza.GetPrice() + 8
}
