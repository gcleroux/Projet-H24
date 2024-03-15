package input

import "github.com/hajimehoshi/ebiten/v2"

type InputHandler interface {
	Update()
	LastPressedInputs() []ebiten.Key
}
