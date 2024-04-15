package events

//
// import (
// 	"github.com/gcleroux/Projet-H24/api/v1"
// 	"github.com/gcleroux/Projet-H24/internal/game/components"
// 	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
// 	"github.com/gcleroux/Projet-H24/internal/game/tags"
// 	"github.com/solarlune/resolv"
// 	"github.com/yohamta/donburi"
// 	"github.com/yohamta/donburi/features/events"
// )
//
// type PeerUpdate struct {
// 	api.PeerPosition
// }
//
// func PeerUpdateHandler(w donburi.World, event PeerUpdate) {
// 	// log.Printf("Getting message from server\n\tID:%d\n\tX:%f\n\tY:%f", event.ID, event.X, event.Y)
//
// 	// TODO: Find a better way to initialize peers
// 	tags.Peer.Each(w, func(e *donburi.Entry) {
// 		peer := components.Peer.Get(e)
// 		if peer.ID == event.ID {
// 			o := dresolv.GetObject(e)
// 			o.Position.X = event.X
// 			o.Position.Y = event.Y
// 			return
// 		}
// 	})
// 	tags.Peer.Each(w, func(e *donburi.Entry) {
// 		peer := components.Peer.Get(e)
// 		if !peer.Present {
// 			// Update the peer
// 			peer.Present = true
// 			peer.ID = event.ID
//
// 			// TODO: We could add "solid" here if we wanted collisions for the players
// 			obj := resolv.NewObject(event.X, event.Y, 16, 24, "solid")
// 			dresolv.SetObject(e, obj)
// 			obj.SetShape(resolv.NewRectangle(0, 0, 16, 24))
// 			return
// 		}
// 	})
// }
//
// var PeerUpdateEvent = events.NewEventType[PeerUpdate]()
