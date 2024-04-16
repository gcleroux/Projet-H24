package connections

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type WSConnection struct {
	ID uuid.UUID

	*websocket.Conn
	Ctx          context.Context
	Closed       bool
	WriteTimeout time.Duration
	ReadTimeout  time.Duration

	mu   sync.Mutex
	logf func(f string, v ...interface{})
}

func (c *WSConnection) Dial(addr string) error {
	conn, _, err := websocket.Dial(c.Ctx, addr, nil)
	if err != nil {
		c.logf("[%s] Dial failed: %v", c.ID.String(), err)
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Closed {
		c.logf("[%s] Dial failed: %v", c.ID.String(), net.ErrClosed)
		return net.ErrClosed
	}
	c.Conn = conn
	c.logf("[%s] Connection established: %s", c.ID, addr)
	return nil
}

func NewWSConnection() *WSConnection {
	wsc := &WSConnection{
		ID: uuid.New(),

		Conn:   &websocket.Conn{},
		Ctx:    context.Background(),
		Closed: false,

		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,

		mu:   sync.Mutex{},
		logf: log.Printf,
	}
	return wsc
}

func (c *WSConnection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Closed = true
	if c.Conn == nil {
		c.logf("[%s] Close failed: %v", c.ID.String(), net.ErrClosed)
		return net.ErrClosed
	}

	if err := c.Conn.CloseNow(); err != nil {
		c.logf("[%s] Close failed: %v", c.ID.String(), err)
		return err
	}

	c.logf("[%s] Connection closed", c.ID.String())
	return nil
}

// func (c *WSConnection) CloseStatus(code websocket.StatusCode, msg string) error {
// 	c.Closed = true
// 	if c.Conn != nil {
// 		if err := c.Conn.Close(websocket.StatusCode(code), msg); err != nil {
// 			c.logf("[%s] Connection closed failed: %v", c.ID.String(), err)
// 			return err
//
// 		}
// 		c.logf("[%s] Connection closed.", c.ID.String())
// 	}
// }

func (c *WSConnection) Write(msg []byte) error {
	// ctx, cancel := context.WithTimeout(c.Ctx, c.WriteTimeout)
	// defer cancel()

	// Sending Marshalled data over the wire
	return c.Conn.Write(c.Ctx, websocket.MessageBinary, msg)
}

func (c *WSConnection) Read() ([]byte, error) {
	// ctx, cancel := context.WithTimeout(c.Ctx, c.ReadTimeout)
	// defer cancel()

	_, data, err := c.Conn.Read(c.Ctx)

	return data, err
}

func (c *WSConnection) Raw() interface{} {
	return c.Conn
}
