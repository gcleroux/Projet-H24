package game

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type rand struct {
	x, y, z, w uint32
}

func (r *rand) next() uint32 {
	// math/rand is too slow to keep 60 FPS on web browsers.
	// Use Xorshift instead: http://en.wikipedia.org/wiki/Xorshift
	t := r.x ^ (r.x << 11)
	r.x, r.y, r.z = r.y, r.z, r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))
	return r.w
}

var theRand = &rand{12345678, 4185243, 776511, 45411}

type Game struct {
	noiseImage *image.RGBA
}

func (g *Game) Update() error {
	// Generate the noise with random RGB values.
	const l = screenWidth * screenHeight
	for i := 0; i < l; i++ {
		x := theRand.next()
		g.noiseImage.Pix[4*i] = uint8(x >> 24)
		g.noiseImage.Pix[4*i+1] = uint8(x >> 16)
		g.noiseImage.Pix[4*i+2] = uint8(x >> 8)
		g.noiseImage.Pix[4*i+3] = 0xff
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.noiseImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Run() error {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
