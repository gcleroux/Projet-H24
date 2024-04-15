package connections

import (
	"context"
	"log"

	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// var (
// 	send_timeout = time.Millisecond * 1000
// 	read_timeout = time.Millisecond * 1000
// )

type WSConnection struct {
	ID   uuid.UUID
	Conn *websocket.Conn
	Ctx  context.Context
	logf func(f string, v ...interface{})
}

func (c *WSConnection) Dial(addr string) {
	conn, _, err := websocket.Dial(c.Ctx, addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.Conn = conn
}

func NewWSConnection() *WSConnection {
	wsc := &WSConnection{
		ID:   uuid.New(),
		Conn: &websocket.Conn{},
		Ctx:  context.Background(),
		logf: log.Printf,
	}
	return wsc
}

func (c *WSConnection) Close() error {
	err := c.Conn.CloseNow()
	if err != nil {
		return err
	}
	return nil
}

func (c *WSConnection) Send(msg api.PlayerPosition) error {
	// ctx, cancel := context.WithTimeout(c.Ctx, send_timeout)
	// defer cancel()

	return wsjson.Write(
		c.Ctx,
		c.Conn,
		msg,
	)
}

func (c *WSConnection) Read() (api.PlayerPosition, error) {
	pp := api.PlayerPosition{}

	// ctx, cancel := context.WithTimeout(ws.Ctx, read_timeout)
	// defer cancel()

	err := wsjson.Read(c.Ctx, c.Conn, &pp)
	if err != nil {
		return api.PlayerPosition{}, err
	}

	return pp, nil
}

func (c *WSConnection) Raw() interface{} {
	return c.Conn
}
