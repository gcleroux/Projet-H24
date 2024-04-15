package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	nw "github.com/gcleroux/Projet-H24/internal/networking/network_server"
	"nhooyr.io/websocket"
)

type gameServer struct {
	serveMux http.ServeMux
	logf     func(f string, v ...interface{})

	connectionHandler *nw.NetworkServer

	// players   map[uuid.UUID]api.PlayerPosition
	// playersMu sync.Mutex
}

func NewGameServer() *gameServer {
	gs := &gameServer{
		logf:              log.Printf,
		connectionHandler: nw.NewNetworkServer(),
		// players:           make(map[uuid.UUID]api.PlayerPosition),
		// playersMu:         sync.Mutex{},
	}

	gs.serveMux.Handle("/ws", http.HandlerFunc(gs.wsHandler))

	return gs
}

func (gs *gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.serveMux.ServeHTTP(w, r)
}

func (gs *gameServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	err := gs.ws(w, r)
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

func (gs *gameServer) ws(w http.ResponseWriter, r *http.Request) error {
	gs.logf("Received /ws call from %s", r.Host)

	conn, err := gs.connectionHandler.Accept(w, r)
	if err != nil {
		return err
	}

	for {
		msg, err := conn.Read()
		if err != nil {
			return err
		}
		gs.connectionHandler.Add(msg.ID, conn)
		gs.connectionHandler.Broadcast(msg)
	}
}
