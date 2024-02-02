package game

import (
	"context"
	"sync"

	"github.com/gcleroux/Projet-H24/internal/server"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type NetworkHandler interface {
	WritePlayerPosition(pos Position) error
	GetPeersPosition() []Position
	Close()
}

// TODO: The network handler shouldn't manage peers, this should be done by the server
type WebSocketNetworkHandler struct {
	id    uuid.UUID
	conn  *websocket.Conn
	peers map[uuid.UUID]Position
	mu    sync.RWMutex
}

func NewWebSocketNetworkHandler(url string) (*WebSocketNetworkHandler, error) {
	conn, _, err := websocket.Dial(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}
	ws := &WebSocketNetworkHandler{
		id:    uuid.New(),
		conn:  conn,
		peers: make(map[uuid.UUID]Position),
		mu:    sync.RWMutex{},
	}
	go ws.readPeersPosition()
	return ws, nil
}

func (ws *WebSocketNetworkHandler) WritePlayerPosition(pos Position) error {
	return wsjson.Write(
		context.Background(),
		ws.conn,
		server.PlayerPositionMessage{
			ClientID: ws.id,
			X:        pos.X,
			Y:        pos.Y,
		},
	)
}

func (ws *WebSocketNetworkHandler) GetPeersPosition() []Position {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	pos := make([]Position, 0, len(ws.peers))
	for _, p := range ws.peers {
		pos = append(pos, p)
	}
	return pos
}

func (ws *WebSocketNetworkHandler) Close() {
	ws.conn.Close(websocket.StatusNormalClosure, "connection closed")
}

func (ws *WebSocketNetworkHandler) readPeersPosition() error {
	var ppm server.PlayerPositionMessage

	for {
		err := wsjson.Read(context.Background(), ws.conn, &ppm)
		if err != nil {
			return err
		}

		// Update the positions of other players
		if ppm.ClientID != ws.id {
			ws.mu.Lock()
			ws.peers[ppm.ClientID] = Position{X: ppm.X, Y: ppm.Y}
			ws.mu.Unlock()
		}
	}
}
