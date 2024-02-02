package main

import (
	"log"

	"github.com/gcleroux/Projet-H24/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(ebiten.RunGame(g))
}
