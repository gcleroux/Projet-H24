package server

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// Player represents a connected player with its WebSocket connection and position.
type Player struct {
	pos       chan position
	closeSlow func()
}

type gameServer struct {
	playerPositionBuffer int
	publishLimiter       *rate.Limiter
	logf                 func(f string, v ...interface{})
	serveMux             http.ServeMux
	playersMu            sync.Mutex
	players              map[*Player]struct{}
}

func NewGameServer(publishRate int) *gameServer {
	gs := &gameServer{
		playerPositionBuffer: 16,
		logf:                 log.Printf,
		players:              make(map[*Player]struct{}),
		publishLimiter: rate.NewLimiter(
			rate.Every(time.Second/time.Duration(publishRate)),
			8,
		),
	}
	gs.serveMux.Handle("/subscribe", http.HandlerFunc(gs.subscribeHandler))
	gs.serveMux.Handle("/publish", http.HandlerFunc(gs.publishHandler))

	return gs
}

func (gs *gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.serveMux.ServeHTTP(w, r)
}

func (gs *gameServer) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := gs.subscribe(r.Context(), w, r)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		gs.logf("%v", err)
		return
	}
}

func (gs *gameServer) publishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket connection", http.StatusInternalServerError)
		return
	}
	defer c.CloseNow()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	playerPosition := position{}

	err = wsjson.Read(ctx, c, &playerPosition)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	gs.publish(playerPosition)

	w.WriteHeader(http.StatusAccepted)
}

func (gs *gameServer) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var mu sync.Mutex
	var c *websocket.Conn
	var closed bool

	p := &Player{
		pos: make(chan position, gs.playerPositionBuffer),
		closeSlow: func() {
			mu.Lock()
			defer mu.Unlock()
			closed = true
			if c != nil {
				c.Close(
					websocket.StatusPolicyViolation,
					"connection too slow to keep up with messages",
				)
			}
		},
	}

	gs.addPlayer(p)
	defer gs.deletePlayer(p)

	c2, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:8080"},
	})
	if err != nil {
		return err
	}
	mu.Lock()
	if closed {
		mu.Unlock()
		return net.ErrClosed
	}
	c = c2
	mu.Unlock()
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)

	for {
		select {
		case pos := <-p.pos:
			err := writeTimeout(ctx, time.Second*5, c, pos)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (gs *gameServer) publish(pos position) {
	gs.playersMu.Lock()
	defer gs.playersMu.Unlock()

	gs.publishLimiter.Wait(context.Background())

	for p := range gs.players {
		select {
		case p.pos <- pos:
		default:
			go p.closeSlow()
		}
	}
}

func (gs *gameServer) addPlayer(p *Player) {
	gs.playersMu.Lock()
	gs.players[p] = struct{}{}
	gs.playersMu.Unlock()
}

func (gs *gameServer) deletePlayer(p *Player) {
	gs.playersMu.Lock()
	delete(gs.players, p)
	gs.playersMu.Unlock()
}

func writeTimeout(
	ctx context.Context,
	timeout time.Duration,
	c *websocket.Conn,
	pos position,
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return wsjson.Write(ctx, c, pos)
}
