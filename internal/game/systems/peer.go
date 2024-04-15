package systems

import (
	"image/color"

	nw "github.com/gcleroux/Projet-H24/internal/networking/network_client"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

// var peerUpd networking.Subscriber[api.PlayerPosition]
// var peerMap map[uuid.UUID]interface{}
//
// func init() {
// 	peerUpd = networking.NewSubscriber[api.PlayerPosition](12)
// 	networking.Connection.AddSubscriber(peerUpd)
// }
//
// func UpdatePeer(ecs *ecs.ECS) {
// 	for len(peerUpd.Chan) > 0 {
// 		msg := <-peerUpd.Chan
// 	// 	ok := false
//
// 	tags.Peer.Each(ecs.World, func(e *donburi.Entry) {
// 		peer := components.Peer.Get(e)
// 		if peer.ID == msg.ID {
// 			// ok = true
// 			o := dresolv.GetObject(e)
// 			o.Position.X = msg.X
// 			o.Position.Y = msg.Y
// 		}
// 	})
// 	// if !ok {
// 	// 	factory.CreatePeer(ecs, msg)
// 	// }
// 	// }
// }

func DrawPeer(_ *ecs.ECS, screen *ebiten.Image) {
	peerColor := color.RGBA{0, 100, 100, 255}
	for _, peer := range nw.NetClient.Peers {
		vector.DrawFilledRect(
			screen,
			float32(peer.X),
			float32(peer.Y),
			float32(16),
			float32(24),
			peerColor,
			false,
		)
	}
}
