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

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	bounds image.Rectangle
	scene  Scene
}

func NewGame() *Game {
	g := &Game{
		bounds: image.Rectangle{},
		scene:  &scenes.PlatformerScene{},
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
	width, height := viper.GetInt("window.width"), viper.GetInt("window.height")
	ebiten.SetWindowSize(width, height)

	// Set window's properties
	if !viper.GetBool("window.resizable") {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	}

	// Start the game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

//
//
//
//
//
//
// const (
// 	screenWidth  = 320
// 	screenHeight = 240
// )
//
// // Generate a new ID for the game client at runtime
// var id uuid.UUID = uuid.New()
//
// // TODO: Refactor the Game struct to use interfaces in all fields
// type Game struct {
// 	id      uuid.UUID
// 	player  *characters.PlayableCharacter
// 	input   input.InputHandler
// 	network networking.NetworkHandler
// }
//
// func NewGame() (*Game, error) {
// 	// Init the game ID
// 	network, err := networking.NewWebSocketNetworkHandler(id)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	g := &Game{
// 		id: id,
// 		player: characters.NewPlayableCharacter(
// 			screenWidth/2,
// 			screenHeight/2,
// 			8,
// 			screenWidth,
// 			screenHeight,
// 		),
// 		input:   &input.KeyboardInputHandler{},
// 		network: network,
// 	}
// 	if err := g.network.SendPlayerPosition(g.player.X, g.player.Y); err != nil {
// 		return nil, err
// 	}
//
// 	return g, nil
// }
//
// func (g *Game) Update() error {
// 	g.input.Update()
//
// 	playerMoved := g.player.Update(g.input.LastPressedInputs())
//
// 	// Send player's position to the server
// 	if playerMoved {
// 		err := g.network.SendPlayerPosition(g.player.X, g.player.Y)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//
// func (g *Game) Draw(screen *ebiten.Image) {
// 	screen.Fill(color.Gray{128})
// 	g.player.Draw(screen)
//
// 	for _, p := range g.network.Peers() {
// 		p.Draw(screen)
// 	}
//
// 	ebitenutil.DebugPrint(
// 		screen,
// 		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
// 	)
// }
//
// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }
//
// func main() {
// 	g, err := game.NewGame()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	log.Fatalln(ebiten.RunGame(g))
// }
