package factories

import (
	"abstract_factory/products/cadeira"
	"abstract_factory/products/mesa"
	"abstract_factory/products/sofa"
)

type VitorianaFactory struct{}

func (f VitorianaFactory) CriarCadeira() cadeira.Cadeira {
	return cadeira.CadeiraVitoriana{}
}

func (f VitorianaFactory) CriarMesa() mesa.Mesa {
	return mesa.MesaVitoriana{}
}

func (f VitorianaFactory) CriarSofa() sofa.Sofa {
	return sofa.SofaVitoriana{}
}
