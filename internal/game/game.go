package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Position struct {
	X float64
	Y float64
}

// TODO: Refactor the Game struct to use interfaces in all fields
type Game struct {
	player  *Player
	input   *InputHandler
	network NetworkHandler
}

func NewGame() (*Game, error) {
	network, err := NewWebSocketNetworkHandler("ws://localhost:8888/position")
	if err != nil {
		return nil, err
	}

	g := &Game{
		player: NewPlayer(screenWidth/2, screenHeight/2, 8),
		input: &InputHandler{
			Keys: []ebiten.Key{},
		},
		network: network,
	}
	if err := g.network.WritePlayerPosition(g.player.Position); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.player.Update(g.input.Keys); err != nil {
		return err
	}

	// Send player's position to the server
	err := g.network.WritePlayerPosition(g.player.Position)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{128})
	g.player.Draw(screen)

	for _, p := range g.network.GetPeersPosition() {
		vector.DrawFilledRect(
			screen,
			float32(p.X),
			float32(p.Y),
			8,
			8,
			color.RGBA{R: 255, G: 0, B: 0, A: 255},
			false,
		)
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
