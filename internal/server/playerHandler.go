package server

import (
	"sync"

	api "github.com/gcleroux/Projet-H24/api/v1"
	"github.com/google/uuid"
)

type PlayerHandler struct {
	// The information about the players on the connected clients
	players   map[uuid.UUID]api.PlayerPosition
	playersMu sync.Mutex
	logf      func(f string, v ...interface{})
}

func (ph *PlayerHandler) addPlayer(id uuid.UUID, ppm api.PlayerPosition) {
	ph.playersMu.Lock()
	ph.players[id] = ppm
	ph.playersMu.Unlock()
}

func (ph *PlayerHandler) deletePlayer(id uuid.UUID, ppm api.PlayerPosition) {
	ph.playersMu.Lock()
	delete(ph.players, id)
	ph.playersMu.Unlock()
}
