package components

import (
	"github.com/yohamta/donburi"
)

type StartCoordsData struct {
	X float64
	Y float64
}

var StartCoords = donburi.NewComponentType[StartCoordsData]()
