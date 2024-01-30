package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"nhooyr.io/websocket"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	playerX float32
	playerY float32
	image   *ebiten.Image
	conn    *websocket.Conn
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += 5
	}

	//
	// // Only send the player position if it has changed
	// if inpututil.IsKeyJustPressed(ebiten.KeyW) ||
	// 	inpututil.IsKeyJustPressed(ebiten.KeyS) ||
	// 	inpututil.IsKeyJustPressed(ebiten.KeyA) ||
	// 	inpututil.IsKeyJustPressed(ebiten.KeyD) {
	//
	// 	// Send player position to the server
	// 	playerPosition := struct {
	// 		X float64 `json:"x"`
	// 		Y float64 `json:"y"`
	// 	}{playerX, playerY}
	//
	// 	if err := conn.WriteJSON(playerPosition); err != nil {
	// 		return err
	// 	}
	// }
	//
	// // Draw the player
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(playerX, playerY)
	// ebitenutil.DebugPrint(screen, "Move the cube with W, A, S, D keys")
	// screen.DrawImage(playerImage, op)
	//
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{128})
	vector.DrawFilledRect(screen, g.playerX, g.playerY, 8.0, 8.0, color.Black, false)

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

	// conn, _, err := websocket.Dial(context.TODO(), "ws://localhost:8080/ws", nil)
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close(websocket.StatusNormalClosure, "")

	g := &Game{
		playerX: screenWidth / 2,
		playerY: screenHeight / 2,
		image:   nil,
		conn:    nil,
	}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
