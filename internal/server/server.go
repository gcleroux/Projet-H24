package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Player represents a player with x and y coordinates.
type Player struct {
	ID uuid.UUID
	Position
}

type gameServer struct {
	serveMux      http.ServeMux
	logf          func(f string, v ...interface{})
	connections   map[*websocket.Conn]struct{}
	connectionsMu sync.Mutex
	players       map[uuid.UUID]Player
	playersMu     sync.Mutex
}

func NewGameServer() *gameServer {
	gs := &gameServer{
		logf:          log.Printf,
		connections:   make(map[*websocket.Conn]struct{}),
		connectionsMu: sync.Mutex{},
		players:       make(map[uuid.UUID]Player),
		playersMu:     sync.Mutex{},
	}
	gs.serveMux.Handle("/position", http.HandlerFunc(gs.positionHandler))

	return gs
}

func (gs *gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.serveMux.ServeHTTP(w, r)
}

func (gs *gameServer) positionHandler(w http.ResponseWriter, r *http.Request) {
	err := gs.position(r.Context(), w, r)
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

func (gs *gameServer) position(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	gs.logf("Received /position call from %s", r.Host)
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:8080"},
	})
	if err != nil {
		return err
	}
	defer conn.CloseNow()

	gs.addConn(conn)
	defer gs.deleteConn(conn)

	for {
		var player Player

		// Read the message from the WebSocket connection.
		err := wsjson.Read(ctx, conn, &player)
		if err != nil {
			gs.logf("[gs.position]: %v", err)
			return err
		}
		// Update the player's position.
		gs.addPlayer(player)
		gs.publish(ctx, player)
	}
}

func (gs *gameServer) publish(ctx context.Context, player Player) {
	for _, conn := range gs.getConns() {
		err := wsjson.Write(ctx, conn, &player)
		if err != nil {
			gs.logf("[gs.publish]: %v", err)
		}
	}
}

func (gs *gameServer) publishPositions() {
	for {
		select {
		case <-time.After(time.Second / 60): // Send updates at a rate of 60 ticks per second.
			gs.broadcast()
		default:
		}
	}
}

func (gs *gameServer) broadcast() {
	ctx := context.Background()

	// Send the positions to each connected client.
	for _, conn := range gs.getConns() {
		for _, player := range gs.getPlayers() {
			err := wsjson.Write(ctx, conn, &player)
			if err != nil {
				gs.logf("%v", err)
			}
		}
	}
}

func (gs *gameServer) addConn(conn *websocket.Conn) {
	gs.connectionsMu.Lock()
	gs.connections[conn] = struct{}{}
	gs.connectionsMu.Unlock()
	gs.logf("Connection added: total %d", len(gs.getConns()))
}

func (gs *gameServer) deleteConn(conn *websocket.Conn) {
	gs.connectionsMu.Lock()
	delete(gs.connections, conn)
	gs.connectionsMu.Unlock()
	gs.logf("Connection removed: total %d", len(gs.getConns()))
}

func (gs *gameServer) addPlayer(p Player) {
	gs.playersMu.Lock()
	gs.players[p.ID] = p
	gs.playersMu.Unlock()
	gs.logf("addPlayer: total %d", len(gs.getPlayers()))
	gs.logf("Player update: %s to [ %0.2f, %0.2f ]", p.ID.String(), p.X, p.Y)
}

func (gs *gameServer) deletePlayer(p Player) {
	gs.playersMu.Lock()
	delete(gs.players, p.ID)
	gs.playersMu.Unlock()
	gs.logf("deletePlayer: total %d", len(gs.getPlayers()))
}

func (gs *gameServer) getConns() []*websocket.Conn {
	gs.connectionsMu.Lock()
	defer gs.connectionsMu.Unlock()

	conns := make([]*websocket.Conn, 0, len(gs.connections))
	for c := range gs.connections {
		conns = append(conns, c)
	}
	return conns
}

func (gs *gameServer) getPlayers() []Player {
	gs.playersMu.Lock()
	defer gs.playersMu.Unlock()

	players := make([]Player, 0, len(gs.players))
	for _, p := range gs.players {
		players = append(players, p)
	}
	return players
}
