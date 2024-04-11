package connections

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type WSConnection struct {
	ID   uuid.UUID
	Conn *websocket.Conn
	logf func(f string, v ...interface{})
}

func NewWebSocketConnection() *WSConnection {
	wsc := &WSConnection{
		ID:   uuid.Nil,
		Conn: &websocket.Conn{},
		logf: log.Printf,
	}
	return wsc
}

func (c *WSConnection) Open(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
		// OriginPatterns: []string{"192.168.0.161:8080"},
	})
	if err != nil {
		return err
	}

	// Assign the connection
	c.ID = uuid.New()
	c.Conn = conn

	return nil
}

func (c *WSConnection) Close() error {
	err := c.Conn.CloseNow()
	if err != nil {
		return err
	}
	return nil
}

func (c *WSConnection) Raw() interface{} {
	return c.Conn
}
