package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/google/uuid"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePeer(ecs *ecs.ECS) *donburi.Entry {
	peer := archetypes.Peer.Spawn(ecs)

	obj := resolv.NewObject(32, 128, 16, 24)
	dresolv.SetObject(peer, obj)

	// Init the peer to not currently in the game
	components.Peer.SetValue(peer, components.PeerData{
		ID:          uuid.Nil,
		Present:     false,
		FacingRight: true,
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 16, 24))

	return peer
}
