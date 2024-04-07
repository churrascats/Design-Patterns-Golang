package game

import "fmt"

var (
	DressFactorySingleInstance = &DressFactory{
		DressMap: make(map[DRESS_TYPE]Dress),
	}
)

type DressFactory struct {
	DressMap map[DRESS_TYPE]Dress
}

func (d *DressFactory) GetDressByType(dressType DRESS_TYPE) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = newTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = newCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return DressFactorySingleInstance
}
