package game

import (
	"context"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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
	// publishConn   *websocket.Conn
	subscribeConn *websocket.Conn
	otherPlayers  map[string]Position
}

func (g *Game) Update() error {
	g.input.Update()
	g.player.Update(g.input.Keys)

	// // Send player's position to the server
	// if err := wsjson.Write(context.Background(), g.publishConn, Position{X: g.player.Position.X, Y: g.player.Position.Y}); err != nil {
	// 	return err
	// }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{128})
	g.player.Draw(screen)

	// Draw other players.
	for _, pos := range g.otherPlayers {
		vector.DrawFilledRect(
			screen,
			float32(pos.X),
			float32(pos.Y),
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

func Run() error {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Multiplayer 2D cube")

	// Initialize the WebSocket connections
	// publishConn, _, err := websocket.Dial(context.Background(), "ws://localhost:8888/publish", nil)
	// if err != nil {
	// return err
	// }

	subscribeConn, _, err := websocket.Dial(
		context.Background(),
		"ws://localhost:8888/subscribe",
		nil,
	)
	if err != nil {
		// publishConn.Close(websocket.StatusNormalClosure, "subscribe connection failed")
		return err
	}

	g := &Game{
		player: NewPlayer(screenWidth/2, screenHeight/2, 8),
		input:  &InputHandler{},
		// publishConn:   publishConn,
		subscribeConn: subscribeConn,
		otherPlayers:  make(map[string]Position),
	}
	go func() {
		for {
			var pos Position
			err := wsjson.Read(context.Background(), subscribeConn, &pos)
			if err != nil {
				// Handle the error (e.g., log it) and break the loop if needed.
				break
			}

			// Update the positions of other players.
			g.otherPlayers["other"] = Position{X: pos.X, Y: pos.Y}
		}
	}()
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	// publishConn.Close(websocket.StatusNormalClosure, "publish closed")
	subscribeConn.Close(websocket.StatusNormalClosure, "subscribe closed")
	return nil
}
