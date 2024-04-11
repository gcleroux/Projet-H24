package networking

import (
	"context"
	"time"

	api "github.com/gcleroux/Projet-H24/api/v1"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Move this to config.yaml
const (
	messageTimeout = time.Second * 1
	idleTimeout    = time.Second * 10
)

type WebSocketClient struct {
	ID   uuid.UUID
	Conn *websocket.Conn
	Ctx  context.Context
}

func NewWebSocketClient() *WebSocketClient {
	return &WebSocketClient{
		ID:   uuid.UUID{},
		Conn: &websocket.Conn{},
		Ctx:  context.Background(),
	}
}

// Opens the connection with the remote server
func (ws *WebSocketClient) Open(addr string) error {
	conn, _, err := websocket.Dial(ws.Ctx, addr, nil)
	if err != nil {
		return err
	}

	ws.Conn = conn
	return nil
}

func (ws *WebSocketClient) Close() error {
	return ws.Conn.Close(websocket.StatusNormalClosure, "connection closed")
}

func (ws *WebSocketClient) Write(pp api.PlayerPosition) error {
	ctx, cancel := context.WithTimeout(ws.Ctx, messageTimeout)
	defer cancel()

	return wsjson.Write(
		ctx,
		ws.Conn,
		pp,
	)
}

func (ws *WebSocketClient) Read() (api.PeerPosition, error) {
	pp := api.PeerPosition{}

	ctx, cancel := context.WithTimeout(ws.Ctx, messageTimeout)
	defer cancel()

	err := wsjson.Read(ctx, ws.Conn, &pp)
	if err != nil {
		return api.PeerPosition{}, err
	}

	return pp, nil
}
