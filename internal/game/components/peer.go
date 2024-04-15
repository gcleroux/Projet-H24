package components

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/gcleroux/Projet-H24/internal/networking"
	"github.com/google/uuid"
	"github.com/yohamta/donburi"
)

type PeerData struct {
	ID          uuid.UUID
	FacingRight bool

	networking.Subscriber[api.PlayerPosition]
}

var Peer = donburi.NewComponentType[PeerData]()
