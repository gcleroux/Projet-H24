package game

import (
	"context"
	"fmt"
	"image/color"
	"log"

	"github.com/gcleroux/Projet-H24/internal/server"
	"github.com/google/uuid"
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
	ID           uuid.UUID
	player       *Player
	input        *InputHandler
	conn         *websocket.Conn
	otherPlayers map[uuid.UUID]Position
}

func (g *Game) Update() error {
	g.input.Update()
	g.player.Update(g.input.Keys)

	// Send player's position to the server
	err := wsjson.Write(
		context.Background(),
		g.conn,
		&server.Player{
			ID:       g.ID,
			Position: server.Position(g.player.Position),
		},
	)
	if err != nil {
		return err
	}
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

	conn, _, err := websocket.Dial(
		context.Background(),
		"ws://localhost:8888/position",
		nil,
	)
	if err != nil {
		return err
	}
	defer conn.Close(websocket.StatusNormalClosure, "subscribe closed")

	g := &Game{
		ID:           uuid.New(),
		player:       NewPlayer(screenWidth/2, screenHeight/2, 8),
		input:        &InputHandler{},
		conn:         conn,
		otherPlayers: make(map[uuid.UUID]Position),
	}

	go func(id uuid.UUID) {
		for {
			var resp server.Player
			err := wsjson.Read(context.Background(), conn, &resp)
			if err != nil {
				log.Printf("%v", err)
				// Handle the error (e.g., log it) and break the loop if needed.
				break
			}
			log.Print("Got position for player: ", resp.ID)

			if resp.ID != id {
				// Update the positions of other players.
				g.otherPlayers[resp.ID] = Position{X: resp.X, Y: resp.Y}
			}

		}
	}(g.ID)

	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
