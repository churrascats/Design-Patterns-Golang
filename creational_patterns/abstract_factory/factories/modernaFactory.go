package factories

import (
	"abstract_factory/products/cadeira"
	"abstract_factory/products/mesa"
	"abstract_factory/products/sofa"
)

type ModernaFactory struct{}

func (f ModernaFactory) CriarCadeira() cadeira.Cadeira {
	return cadeira.CadeiraModerna{}
}

func (f ModernaFactory) CriarMesa() mesa.Mesa {
	return mesa.MesaModerna{}
}

func (f ModernaFactory) CriarSofa() sofa.Sofa {
	return sofa.SofaModerna{}
}
