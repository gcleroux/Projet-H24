package factory

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePeer(ecs *ecs.ECS, pp api.PlayerPosition) *donburi.Entry {
	peer := archetypes.Peer.Spawn(ecs)

	obj := resolv.NewObject(pp.X, pp.Y, 16, 24)
	dresolv.SetObject(peer, obj)
	obj.SetShape(resolv.NewRectangle(0, 0, 16, 24))

	// sub := networking.NewSubscriber[api.PlayerPosition](100)
	// networking.Connection.AddSubscriber(sub)

	// Init the peer to not currently in the game
	components.Peer.SetValue(peer, components.PeerData{
		ID:          pp.ID,
		FacingRight: true,
	})

	return peer
}
