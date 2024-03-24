package factories

import (
	"abstract_factory/products"
	"fmt"
)

type FactoryManager struct{}

func (f *FactoryManager) GetFactory(factoryType products.FactoryType) AbstractFactory {
	var factory AbstractFactory

	switch factoryType {

	case products.Moderna:
		factory = ModernaFactory{}

	case products.Vitoriana:
		factory = VitorianaFactory{}

	case products.Original:
		factory = OriginalFactory{}

	default:
		fmt.Println("Factory not found")
	}

	return factory
}
