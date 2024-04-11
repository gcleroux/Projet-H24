package components

import (
	"github.com/google/uuid"
	"github.com/yohamta/donburi"
)

type PeerData struct {
	ID          uuid.UUID
	Present     bool
	FacingRight bool
}

var Peer = donburi.NewComponentType[PeerData]()
