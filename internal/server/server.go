package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gcleroux/Projet-H24/internal/server/connections"
	"github.com/gcleroux/Projet-H24/internal/server/messages"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type gameServer struct {
	serveMux http.ServeMux
	logf     func(f string, v ...interface{})

	connectionHandler connections.ConnectionHandler

	players   map[uuid.UUID]messages.PlayerPositionMessage
	playersMu sync.Mutex
}

func NewGameServer() *gameServer {
	gs := &gameServer{
		logf:              log.Printf,
		connectionHandler: connections.NewWSConnectionHandler(),
		players:           make(map[uuid.UUID]messages.PlayerPositionMessage),
		playersMu:         sync.Mutex{},
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

	// Create and open the websocket connection
	conn := connections.NewWebSocketConnection()
	err := conn.Open(r.Context(), w, r)
	if err != nil {
		gs.logf("%v", err)
		return err
	}

	gs.connectionHandler.Add(conn)
	defer gs.connectionHandler.Remove(conn)

	for {
		var ppm messages.PlayerPositionMessage

		// Read the message from the WebSocket connection.
		err := wsjson.Read(ctx, conn.Conn, &ppm)
		if err != nil {
			gs.logf("[gs.position]: %v", err)
			return err
		}
		// Update the player's position.
		gs.addPlayer(conn.ClientID, ppm)
		gs.publish(ctx, ppm)
	}
}

func (gs *gameServer) publish(ctx context.Context, ppm messages.PlayerPositionMessage) {
	for _, conn := range gs.connectionHandler.GetConns() {
		if c, ok := conn.(*connections.WSConnection); ok {
			err := wsjson.Write(ctx, c.Conn, ppm)
			if err != nil {
				gs.logf("[gs.publish]: %v", err)
			}
		}
	}
}

func (gs *gameServer) addPlayer(id uuid.UUID, ppm messages.PlayerPositionMessage) {
	gs.playersMu.Lock()
	gs.players[id] = ppm
	gs.playersMu.Unlock()
}
