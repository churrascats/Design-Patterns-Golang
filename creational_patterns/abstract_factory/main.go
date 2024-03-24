package main

import (
	"abstract_factory/factories"
	"abstract_factory/products"
)

func main() {
	factoryManager := factories.FactoryManager{}

	factory := factoryManager.GetFactory(products.Vitoriana)

	cadeira := factory.CriarCadeira()
	mesa := factory.CriarMesa()
	sofa := factory.CriarSofa()

	cadeira.Sentar()
	mesa.ColocarObjeto()
	sofa.Encostar()
}
