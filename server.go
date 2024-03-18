//go:build server
// +build server

package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gcleroux/Projet-H24/internal/server"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on http://%v", l.Addr())

	gs := server.NewGameServer()
	s := &http.Server{
		Handler: gs,
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
