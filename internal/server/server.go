package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type gameServer struct {
	serveMux http.ServeMux
	logf     func(f string, v ...interface{})

	connections   map[*websocket.Conn]clientID
	connectionsMu sync.Mutex

	players   map[clientID]PlayerPositionMessage
	playersMu sync.Mutex
}

func NewGameServer() *gameServer {
	gs := &gameServer{
		logf:          log.Printf,
		connections:   make(map[*websocket.Conn]clientID),
		connectionsMu: sync.Mutex{},
		players:       make(map[clientID]PlayerPositionMessage),
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
		InsecureSkipVerify: true,
		// OriginPatterns: []string{"192.168.0.161:8080"},
	})
	if err != nil {
		return err
	}
	defer conn.CloseNow()

	// Read the registration message from the WebSocket connection.
	var registration RegisterClientMessage
	err = wsjson.Read(ctx, conn, &registration)
	if err != nil {
		gs.logf("[gs.position]: %v", err)
		return err
	}

	// Use the game ID from the registration message to uniquely identify the game instance.
	clientID := registration.ClientID
	gs.logf("Game registered with ID: %s", clientID.String())

	gs.addConn(conn, clientID)
	defer gs.deleteConn(conn)

	for {
		var ppm PlayerPositionMessage

		// Read the message from the WebSocket connection.
		err := wsjson.Read(ctx, conn, &ppm)
		if err != nil {
			gs.logf("[gs.position]: %v", err)
			return err
		}
		// Update the player's position.
		gs.addPlayer(clientID, ppm)
		gs.publish(ctx, ppm)
	}
}

func (gs *gameServer) publish(ctx context.Context, ppm PlayerPositionMessage) {
	for _, conn := range gs.getConns() {
		err := wsjson.Write(ctx, conn, ppm)
		if err != nil {
			gs.logf("[gs.publish]: %v", err)
		}
	}
}

func (gs *gameServer) addConn(conn *websocket.Conn, id clientID) {
	gs.connectionsMu.Lock()
	gs.connections[conn] = id
	gs.connectionsMu.Unlock()
	gs.logf("Connection added: total %d", len(gs.getConns()))
}

func (gs *gameServer) deleteConn(conn *websocket.Conn) {
	gs.connectionsMu.Lock()
	delete(gs.connections, conn)
	gs.connectionsMu.Unlock()
	gs.logf("Connection removed: total %d", len(gs.getConns()))
}

func (gs *gameServer) addPlayer(id clientID, ppm PlayerPositionMessage) {
	gs.playersMu.Lock()
	gs.players[id] = ppm
	gs.playersMu.Unlock()
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
