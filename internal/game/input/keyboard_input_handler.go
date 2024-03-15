package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// KeyboardInputHandler implements the InputHandler interface.
// It listens for pressed keyboard keys at every new frame and
// appends them to the Keys buffer.
type KeyboardInputHandler struct {
	Keys []ebiten.Key
}

// Saves all the pressed keys in a buffer
// This function is called every frame
func (kb *KeyboardInputHandler) Update() {
	kb.Keys = inpututil.AppendPressedKeys(kb.Keys[:0])
}

// Returns all the pressed keys in the last frame
func (kb KeyboardInputHandler) LastPressedInputs() []ebiten.Key {
	return kb.Keys
}
