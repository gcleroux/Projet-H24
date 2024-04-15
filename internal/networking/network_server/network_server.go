package network_server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gcleroux/Projet-H24/api/v1"
	co "github.com/gcleroux/Projet-H24/internal/networking/connections"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

// A NetworkServer keeps track of the clients connected to the
// Game room using Websockets
type NetworkServer struct {
	connections   map[uuid.UUID]co.Connection
	connectionsMu sync.Mutex
	logf          func(f string, v ...interface{})
}

func NewNetworkServer() *NetworkServer {
	n := &NetworkServer{
		connections:   make(map[uuid.UUID]co.Connection),
		connectionsMu: sync.Mutex{},
		logf:          log.Printf,
	}
	return n
}

func (n *NetworkServer) Accept(
	w http.ResponseWriter,
	r *http.Request,
) (co.Connection, error) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
		// OriginPatterns: []string{"192.168.0.161:8080"},
	})
	if err != nil {
		return &co.WSConnection{}, err
	}

	return &co.WSConnection{Conn: conn, Ctx: r.Context()}, nil
}

func (n *NetworkServer) Add(id uuid.UUID, conn co.Connection) {
	n.connectionsMu.Lock()
	n.connections[id] = conn
	n.connectionsMu.Unlock()
}

func (n *NetworkServer) Remove(id uuid.UUID) {
	n.connectionsMu.Lock()
	delete(n.connections, id)
	n.connectionsMu.Unlock()

	n.logf("Connection removed: total %d", len(n.GetConns()))
}

func (n *NetworkServer) GetConns() []co.Connection {
	n.connectionsMu.Lock()
	defer n.connectionsMu.Unlock()

	conns := make([]co.Connection, 0, len(n.connections))
	for _, c := range n.connections {
		conns = append(conns, c)
	}
	return conns
}

func (n *NetworkServer) Broadcast(msg api.PlayerPosition) {
	n.connectionsMu.Lock()
	defer n.connectionsMu.Unlock()

	for id, conn := range n.connections {
		if id != msg.ID {
			msg.ServerT = time.Now().UnixMilli()
			go conn.Send(msg)
		}
	}
}
