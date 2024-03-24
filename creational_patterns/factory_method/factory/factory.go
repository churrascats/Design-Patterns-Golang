package factory

import "fmt"

type PolygonFactory struct {
}

func (p *PolygonFactory) GetPolygon(polygonType PolygonType) Polygon {
	var polygon Polygon

	switch polygonType {
	case CIRCLE:
		polygon = &Circle{}

	case SQUARE:
		polygon = &Square{}

	case TRIANGLE:
		polygon = &Triangle{}

	default:
		fmt.Println("Polygon doesn't exists")
	}

	return polygon
}
