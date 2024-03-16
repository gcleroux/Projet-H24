package characters

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Data structure that reprensents the current player
type PlayableCharacter struct {
	Image                    *ebiten.Image
	X                        float64
	Y                        float64
	Speed                    float64
	Velocity                 float64
	UpperBoundX, UpperBoundY float64
}

// Create a new player for the game at position x,y.
// Size is the size of the player sprite in pixels
func NewPlayableCharacter(
	x, y float64,
	size int,
	screenWidth, screenHeigth int,
) *PlayableCharacter {
	p := &PlayableCharacter{
		Image:       ebiten.NewImage(size, size),
		X:           x,
		Y:           y,
		Speed:       5.0,
		Velocity:    0.0,
		UpperBoundX: float64(screenWidth - size),
		UpperBoundY: float64(screenHeigth - size),
	}
	p.Image.Fill(color.Black)

	return p
}

// Update updates the player's position based on user input
// Controls with WASD, player velocity is normalized
func (p *PlayableCharacter) Update(inputs []ebiten.Key) bool {
	// Player's (x, y) position delta
	// (0, 0) means the player didn't move
	x := 0.0
	y := 0.0

	for _, key := range inputs {
		switch key {
		case ebiten.KeyW:
			y--
		case ebiten.KeyS:
			y++
		case ebiten.KeyA:
			x--
		case ebiten.KeyD:
			x++
		}
	}

	length := math.Abs(x) + math.Abs(y)

	switch {
	case length == 0:
		return false
	case length == 1:
		p.Velocity = p.Speed
	case length == 2:
		p.Velocity = p.Speed / math.Sqrt2
	}

	p.X += (x * p.Velocity)
	p.Y += (y * p.Velocity)

	p.X = max(0, min(p.X, p.UpperBoundX))
	p.Y = max(0, min(p.Y, p.UpperBoundY))

	return true
}

// Draw the player sprite on screen at his position
// Currently, the player is just a square
func (p *PlayableCharacter) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(p.Image, op)
}
