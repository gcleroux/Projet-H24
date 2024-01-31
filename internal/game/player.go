package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Image                    *ebiten.Image
	Position                 Position
	Speed                    float64
	Velocity                 float64
	UpperBoundX, UpperBoundY float64
}

func NewPlayer(x, y float64, size int) *Player {
	p := &Player{
		Image: ebiten.NewImage(size, size),
		Position: Position{
			X: x,
			Y: y,
		},
		Speed:       5.0,
		Velocity:    0.0,
		UpperBoundX: float64(screenWidth - size),
		UpperBoundY: float64(screenHeight - size),
	}
	p.Image.Fill(color.Black)

	return p
}

// Update updates the player's position based on user input
func (p *Player) Update(inputs []ebiten.Key) error {
	dir := Position{X: 0.0, Y: 0.0}

	for _, key := range inputs {
		switch key {
		case ebiten.KeyW:
			dir.Y--
		case ebiten.KeyS:
			dir.Y++
		case ebiten.KeyA:
			dir.X--
		case ebiten.KeyD:
			dir.X++
		}
	}

	length := math.Abs(dir.X) + math.Abs(dir.Y)

	switch {
	case length == 0:
		return nil
	case length == 1:
		p.Velocity = p.Speed
	case length == 2:
		p.Velocity = p.Speed / math.Sqrt2
	}

	p.Position.X += (dir.X * p.Velocity)
	p.Position.Y += (dir.Y * p.Velocity)

	p.Position.X = max(0, min(p.Position.X, p.UpperBoundX))
	p.Position.Y = max(0, min(p.Position.Y, p.UpperBoundY))

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.Image, op)
}
