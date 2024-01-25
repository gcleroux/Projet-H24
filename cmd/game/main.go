package main

import (
	"log"

	"github.com/gcleroux/Projet-H24/internal/game"
)

func main() {
	if err := game.Run(); err != nil {
		log.Fatalln(err)
	}
}
