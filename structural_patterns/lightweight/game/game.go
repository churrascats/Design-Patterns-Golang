package game

type Game struct {
	Terrorists        []*Player
	CounterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		Terrorists:        make([]*Player, 1),
		CounterTerrorists: make([]*Player, 1),
	}
}

func (c *Game) AddTerrorist(dressType DRESS_TYPE) {
	player := newPlayer("T", dressType)
	c.Terrorists = append(c.Terrorists, player)
}

func (c *Game) AddCounterTerrorist(dressType DRESS_TYPE) {
	player := newPlayer("CT", dressType)
	c.CounterTerrorists = append(c.CounterTerrorists, player)
}
