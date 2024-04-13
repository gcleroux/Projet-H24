//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/gcleroux/Projet-H24/internal/game/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var config []byte

func init() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	bounds image.Rectangle
	scene  scenes.Scene
}

func NewGame() *Game {
	g := &Game{
		bounds: image.Rectangle{},
		scene:  scenes.NewSceneManager(),
	}

	return g
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	// Set the window dimensions
	width, height := viper.GetInt(
		"game.settings.window.width",
	), viper.GetInt(
		"game.settings.window.height",
	)
	ebiten.SetWindowSize(width, height)

	// Set window's properties
	if !viper.GetBool("game.settings.window.resizable") {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	}

	// Start the game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
