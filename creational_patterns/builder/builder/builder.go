package builder

import "builder/pizza"

type Builder interface {
	SetSize(int)
	SetCheese(bool)
	SetSauce(bool)
	SetSausage(bool)
	SetBacon(bool)
	Build() pizza.IPizza
}
