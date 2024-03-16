package server

import (
	"github.com/gcleroux/Projet-H24/internal/server/connections"
	"github.com/google/uuid"
)

// A GameRoom is an instance of a multiplayer
// game with at least 1 client connected
type GameRoom struct {
	id uuid.UUID

	Title    string
	Size     int
	Capacity int

	// The clients connected to the lobby
	ConnectionHandler *connections.ConnectionHandler
}
