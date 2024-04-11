package events

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/yohamta/donburi/features/events"
)

type PlayerUpdate struct {
	api.PlayerPosition
}

var PlayerUpdateEvent = events.NewEventType[PlayerUpdate]()
