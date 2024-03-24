package main

import (
	"decorator/decorators"
	"decorator/pizza"
	"fmt"
)

func main() {
	veggiePizza := &pizza.VeggiePizza{}

	calabresa := &decorators.CalabresaDecorator{Pizza: veggiePizza}

	catupiry := &decorators.CatupiryDecorator{Pizza: calabresa}

	frango := &decorators.FrangoDecorator{Pizza: catupiry}

	fmt.Println(frango.GetPrice())
}
