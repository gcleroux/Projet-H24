package systems

import (
	"log"

	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/events"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/yohamta/donburi"
)

func PeerUpdateHandler(w donburi.World, event events.PeerUpdate) {
	log.Printf("Getting message from server\n\tID:%d\n\tX:%f\n\tY:%f", event.ID, event.X, event.Y)
	// log.Print("Getting playerPos")
	tags.Peer.Each(w, func(e *donburi.Entry) {
		peer := components.Peer.Get(e)
		if peer.ID == event.ID {
			o := dresolv.GetObject(e)
			o.Position.X = event.X
			o.Position.Y = event.Y
			return
		}
	})
	tags.Peer.Each(w, func(e *donburi.Entry) {
		peer := components.Peer.Get(e)
		if !peer.Present {
			// Update the peer
			peer.Present = true
			peer.ID = event.ID

			o := dresolv.GetObject(e)
			o.Position.X = event.X
			o.Position.Y = event.Y
			return
		}
	})
}

func PlayerUpdateHandler(w donburi.World, event events.PlayerUpdate) {
	log.Printf("Sending message to server\n\tX:%f\n\tY:%f", event.X, event.Y)
	// log.Print("Sending playerPos")
	entry, _ := components.Connection.First(w)
	if err := components.Connection.Get(entry).Write(event.PlayerPosition); err != nil {
		// Handle error
		panic(err)
	}
}
