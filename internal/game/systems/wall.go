package systems

import (
	"image/color"

	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawWall(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Wall.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{60, 60, 60, 255}
		vector.DrawFilledRect(
			screen,
			float32(o.Position.X),
			float32(o.Position.Y),
			float32(o.Size.X),
			float32(o.Size.Y),
			drawColor,
			false,
		)
	})
}
