package api

import "github.com/google/uuid"

type PlayerPosition struct {
	ID uuid.UUID `json:"id,omitempty"`
	X  float64   `json:"x,omitempty"`
	Y  float64   `json:"y,omitempty"`

	// Timestamps to calculate network latency
	ClientT int64 `json:"client_t,omitempty"`
	ServerT int64 `json:"server_t,omitempty"`
	TotalT  int64 `json:"total_t,omitempty"`
}

type PeerPosition struct {
	ID uuid.UUID `json:"id,omitempty"`
	X  float64   `json:"x,omitempty"`
	Y  float64   `json:"y,omitempty"`
}
