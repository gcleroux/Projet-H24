package events

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/yohamta/donburi/features/events"
)

type PeerUpdate struct {
	api.PeerPosition
}

var PeerUpdateEvent = events.NewEventType[PeerUpdate]()
