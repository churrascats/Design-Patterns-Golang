package main

import "builder/builder"

func main() {
	pizzaBuilder := builder.PizzaBuilder{}

	pizzaBuilder.SetSize(20)
	pizzaBuilder.SetBacon(true)
	pizzaBuilder.SetCheese(true)

	pizza := pizzaBuilder.Build()

	pizza.PrintRecipe()
}
