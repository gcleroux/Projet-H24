package server

import (
	"sync"

	"github.com/gcleroux/Projet-H24/internal/server/messages"
	"github.com/google/uuid"
)

type PlayerHandler struct {
	// The information about the players on the connected clients
	players   map[uuid.UUID]messages.PlayerPositionMessage
	playersMu sync.Mutex
	logf      func(f string, v ...interface{})
}

func (ph *PlayerHandler) addPlayer(id uuid.UUID, ppm messages.PlayerPositionMessage) {
	ph.playersMu.Lock()
	ph.players[id] = ppm
	ph.playersMu.Unlock()
}

func (ph *PlayerHandler) deletePlayer(id uuid.UUID, ppm messages.PlayerPositionMessage) {
	ph.playersMu.Lock()
	delete(ph.players, id)
	ph.playersMu.Unlock()
}
