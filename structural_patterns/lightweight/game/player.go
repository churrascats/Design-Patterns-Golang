package game

type Player struct {
	Dress      Dress
	PlayerType string
	Lat        int
	Long       int
}

func newPlayer(playerType string, dressType DRESS_TYPE) *Player {
	dress, err := GetDressFactorySingleInstance().GetDressByType(dressType)
	if err != nil {
		panic(err)
	}

	return &Player{
		PlayerType: playerType,
		Dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.Lat = lat
	p.Long = long
}
