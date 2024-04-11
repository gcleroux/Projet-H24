package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/google/uuid"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePeer(ecs *ecs.ECS) *donburi.Entry {
	peer := archetypes.Peer.Spawn(ecs)

	// Init the peer to not currently in the game
	components.Peer.SetValue(peer, components.PeerData{
		ID:          uuid.Nil,
		Present:     false,
		FacingRight: true,
	})

	return peer
}
