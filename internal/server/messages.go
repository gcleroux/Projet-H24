package server

import "github.com/google/uuid"

type clientID = uuid.UUID

type RegisterClientMessage struct {
	ClientID clientID `json:"client_id,omitempty"`
}

type PlayerPositionMessage struct {
	ClientID clientID `json:"client_id,omitempty"`
	X        float64  `json:"x,omitempty"`
	Y        float64  `json:"y,omitempty"`
}
