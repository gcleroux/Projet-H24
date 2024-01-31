package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"nhooyr.io/websocket"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Position struct {
	X float64
	Y float64
}

type Game struct {
	player *Player
	input  *InputHandler
	conn   *websocket.Conn
}

func (g *Game) Update() error {
	g.input.Update()
	g.player.Update(g.input.Keys)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{128})
	g.player.Draw(screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Run() error {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Multiplayer 2D cube")

	g := &Game{
		player: NewPlayer(screenWidth/2, screenHeight/2, 8),
		input:  &InputHandler{},
		conn:   nil,
	}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
