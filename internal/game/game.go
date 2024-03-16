package game

import (
	"fmt"
	"image/color"

	"github.com/gcleroux/Projet-H24/internal/game/characters"
	"github.com/gcleroux/Projet-H24/internal/game/input"
	"github.com/gcleroux/Projet-H24/internal/game/networking"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

// Generate a new ID for the game client at runtime
var id uuid.UUID = uuid.New()

// TODO: Refactor the Game struct to use interfaces in all fields
type Game struct {
	id      uuid.UUID
	player  *characters.PlayableCharacter
	input   input.InputHandler
	network networking.NetworkHandler
}

func NewGame() (*Game, error) {
	// Init the game ID
	network, err := networking.NewWebSocketNetworkHandler(id)
	if err != nil {
		return nil, err
	}

	g := &Game{
		id: id,
		player: characters.NewPlayableCharacter(
			screenWidth/2,
			screenHeight/2,
			8,
			screenWidth,
			screenHeight,
		),
		input:   &input.KeyboardInputHandler{},
		network: network,
	}
	if err := g.network.SendPlayerPosition(g.player.X, g.player.Y); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()

	playerMoved := g.player.Update(g.input.LastPressedInputs())

	// Send player's position to the server
	if playerMoved {
		err := g.network.SendPlayerPosition(g.player.X, g.player.Y)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{128})
	g.player.Draw(screen)

	for _, p := range g.network.Peers() {
		p.Draw(screen)
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
