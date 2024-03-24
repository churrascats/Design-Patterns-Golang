package decorators

import "decorator/pizza"

type CalabresaDecorator struct {
	Pizza pizza.Pizza
}

func (c *CalabresaDecorator) GetPrice() int {
	return c.Pizza.GetPrice() + 10
}
