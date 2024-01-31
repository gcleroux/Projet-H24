package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputHandler struct {
	Keys []ebiten.Key
}

func (i *InputHandler) Update() {
	i.Keys = inpututil.AppendPressedKeys(i.Keys[:0])
}
