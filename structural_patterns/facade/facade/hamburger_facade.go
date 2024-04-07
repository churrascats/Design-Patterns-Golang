package facade

import hamburguerplace "facade/hamburguer_place"

type HamburguerFacade struct {
	cheese  *hamburguerplace.Cheese
	lettuce *hamburguerplace.Letucce
	tomato  *hamburguerplace.Tomato
	patty   *hamburguerplace.Patty
}

func NewHamburguerFacade() *HamburguerFacade {
	return &HamburguerFacade{
		cheese:  hamburguerplace.NewCheese(),
		lettuce: hamburguerplace.NewLetucce(),
		tomato:  hamburguerplace.NewTomato(),
		patty:   hamburguerplace.NewPatty(),
	}
}

func (hf *HamburguerFacade) GetHamburguer() string {
	hamburguer := "bun"

	hamburguer += hf.cheese.AddCheese()
	hamburguer += hf.tomato.AddTomato()
	hamburguer += hf.lettuce.AddLetucce()
	hamburguer += hf.patty.AddPatty()

	return hamburguer
}
