package api

import "github.com/google/uuid"

type PlayerPosition struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type PeerPosition struct {
	ID uuid.UUID `json:"id,omitempty"`
	X  float64   `json:"x,omitempty"`
	Y  float64   `json:"y,omitempty"`
}
