package connections

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

// A WSConnectionHandler keeps track of the clients connected to the
// Game room using Websockets
type WSConnectionHandler struct {
	connections   map[Connection]uuid.UUID
	connectionsMu sync.Mutex
	logf          func(f string, v ...interface{})
}

func NewWSConnectionHandler() *WSConnectionHandler {
	ws := &WSConnectionHandler{
		connections:   make(map[Connection]uuid.UUID),
		connectionsMu: sync.Mutex{},
		logf:          log.Printf,
	}
	return ws
}

func (ws *WSConnectionHandler) Add(conn Connection) {
	if c, ok := conn.(*WSConnection); ok {
		ws.connectionsMu.Lock()
		ws.connections[c] = c.ID
		ws.connectionsMu.Unlock()

		ws.logf("Connection added: total %d", len(ws.GetConns()))
	}
}

func (ws *WSConnectionHandler) Remove(conn Connection) {
	ws.connectionsMu.Lock()
	delete(ws.connections, conn)
	ws.connectionsMu.Unlock()

	ws.logf("Connection removed: total %d", len(ws.GetConns()))
}

func (ws *WSConnectionHandler) GetConns() []Connection {
	ws.connectionsMu.Lock()
	defer ws.connectionsMu.Unlock()

	conns := make([]Connection, 0, len(ws.connections))
	for c := range ws.connections {
		conns = append(conns, c)
	}
	return conns
}
