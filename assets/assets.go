package assets

import (
	_ "embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed DejaVuSans.ttf
var fontData []byte

var DefaultFont font.Face

func init() {
	f, err := opentype.Parse(fontData)
	if err != nil {
		panic(err)
	}

	DefaultFont, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    20.0,
		DPI:     96,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}
