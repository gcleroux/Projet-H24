package networking

import (
	"context"
	"sync"

	"github.com/gcleroux/Projet-H24/config"
	"github.com/gcleroux/Projet-H24/internal/game/characters"
	"github.com/gcleroux/Projet-H24/internal/server/messages"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WebSocketNetworkHandler struct {
	id    uuid.UUID
	conn  *websocket.Conn
	peers map[uuid.UUID]characters.Peer
	mu    sync.RWMutex
}

func NewWebSocketNetworkHandler(id uuid.UUID) (*WebSocketNetworkHandler, error) {
	ws := &WebSocketNetworkHandler{
		id:    id,
		conn:  &websocket.Conn{},
		peers: make(map[uuid.UUID]characters.Peer),
		mu:    sync.RWMutex{},
	}

	// Attempts to open the connection with the remote server
	err := ws.Open()
	if err != nil {
		return nil, err
	}

	go ws.ReadPeerPosition()
	return ws, nil
}

// Opens the connection with the remote server
func (ws *WebSocketNetworkHandler) Open() error {
	conn, _, err := websocket.Dial(context.Background(), config.SERVER_URL+"/position", nil)
	if err != nil {
		return err
	}

	ws.conn = conn
	return nil
}

func (ws *WebSocketNetworkHandler) Close() error {
	return ws.conn.Close(websocket.StatusNormalClosure, "connection closed")
}

func (ws *WebSocketNetworkHandler) SendPlayerPosition(x, y float64) error {
	return wsjson.Write(
		context.Background(),
		ws.conn,
		messages.PlayerPositionMessage{
			ClientID: ws.id,
			X:        x,
			Y:        y,
		},
	)
}

func (ws *WebSocketNetworkHandler) ReadPeerPosition() error {
	var ppm messages.PlayerPositionMessage

	for {
		err := wsjson.Read(context.Background(), ws.conn, &ppm)
		if err != nil {
			return err
		}
		go ws.UpdatePeer(ppm)

	}
}

func (ws *WebSocketNetworkHandler) UpdatePeer(ppm messages.PlayerPositionMessage) {
	// Update the positions of other players
	if ppm.ClientID != ws.id {
		ws.mu.RLock()
		peer, ok := ws.peers[ppm.ClientID]
		ws.mu.RUnlock()

		if ok == false {
			peer = *characters.NewPeer(ppm.X, ppm.Y)
		} else {
			peer.X = ppm.X
			peer.Y = ppm.Y
		}
		ws.mu.Lock()
		ws.peers[ppm.ClientID] = peer
		ws.mu.Unlock()
	}
}

func (ws *WebSocketNetworkHandler) Peers() []characters.Peer {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	pos := make([]characters.Peer, 0, len(ws.peers))
	for _, p := range ws.peers {
		pos = append(pos, p)
	}
	return pos
}
