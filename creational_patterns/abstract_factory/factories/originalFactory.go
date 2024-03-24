package factories

import (
	"abstract_factory/products/cadeira"
	"abstract_factory/products/mesa"
	"abstract_factory/products/sofa"
)

type OriginalFactory struct{}

func (f OriginalFactory) CriarCadeira() cadeira.Cadeira {
	return cadeira.CadeiraOriginal{}
}

func (f OriginalFactory) CriarMesa() mesa.Mesa {
	return mesa.MesaOriginal{}
}

func (f OriginalFactory) CriarSofa() sofa.Sofa {
	return sofa.SofaOriginal{}
}
