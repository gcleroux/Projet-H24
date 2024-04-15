package systems

import (
	"image/color"

	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawRamp(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Ramp.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{255, 50, 100, 255}
		tri := o.Shape.(*resolv.ConvexPolygon)
		drawPolygon(screen, tri, drawColor)
	})
}

func drawPolygon(screen *ebiten.Image, polygon *resolv.ConvexPolygon, color color.Color) {
	for _, line := range polygon.Lines() {
		vector.StrokeLine(
			screen,
			float32(line.Start.X),
			float32(line.Start.Y),
			float32(line.End.X),
			float32(line.End.Y),
			1.0,
			color,
			false,
		)
	}
}
