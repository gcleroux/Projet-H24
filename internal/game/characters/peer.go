package characters

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Peer struct {
	X      float64
	Y      float64
	Width  float32
	Height float32
	Color  color.Color
}

func NewPeer(x, y float64) *Peer {
	p := &Peer{
		X:      x,
		Y:      y,
		Width:  8.0,
		Height: 8.0,

		// Set a random color for the peer
		Color: color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		},
	}
	return p
}

func (p *Peer) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(p.X),
		float32(p.Y),
		float32(p.Width),
		float32(p.Height),
		p.Color,
		false,
	)
}
