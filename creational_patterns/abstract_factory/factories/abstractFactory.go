package factories

import (
	"abstract_factory/products/cadeira"
	"abstract_factory/products/mesa"
	"abstract_factory/products/sofa"
)

type AbstractFactory interface {
	CriarCadeira() cadeira.Cadeira
	CriarMesa() mesa.Mesa
	CriarSofa() sofa.Sofa
}
