package events

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
)

type PlayerUpdate struct {
	api.PlayerPosition
}

func PlayerUpdateHandler(w donburi.World, event PlayerUpdate) {
	// log.Printf("Sending message to server\n\tX:%f\n\tY:%f", event.X, event.Y)
	entry := components.Connection.MustFirst(w)
	if err := components.Connection.Get(entry).Write(event.PlayerPosition); err != nil {
		// Handle error
		panic(err)
	}
}

var PlayerUpdateEvent = events.NewEventType[PlayerUpdate]()
