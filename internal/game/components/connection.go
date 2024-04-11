package components

import (
	"github.com/gcleroux/Projet-H24/internal/networking"
	"github.com/yohamta/donburi"
)

type ConnectionData struct {
	*networking.WebSocketClient
}

var Connection = donburi.NewComponentType[ConnectionData]()
