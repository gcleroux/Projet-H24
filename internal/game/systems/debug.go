package systems

import (
	"fmt"
	"image/color"

	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func DrawDebug(ecs *ecs.ECS, screen *ebiten.Image) {
	settings := GetOrCreateSettings(ecs)
	if !settings.Debug {
		return
	}
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
	)

	playerEntry, ok := tags.Player.First(ecs.World)
	if !ok {
		return
	}

	player := components.Player.Get(playerEntry)
	o := dresolv.GetObject(playerEntry)

	// We draw the player as a different color when jumping so we can visually see when he's in the air.
	if player.OnGround == nil {
		vector.DrawFilledRect(
			screen,
			float32(o.Position.X),
			float32(o.Position.Y),
			float32(o.Size.X),
			float32(o.Size.Y),
			color.RGBA{200, 0, 200, 255},
			false,
		)
	}

	spaceEntry, ok := components.Space.First(ecs.World)
	if !ok {
		return
	}
	space := components.Space.Get(spaceEntry)

	for y := 0; y < space.Height(); y++ {
		for x := 0; x < space.Width(); x++ {

			cell := space.Cell(x, y)

			cw := float32(space.CellWidth)
			ch := float32(space.CellHeight)
			cx := float32(cell.X) * cw
			cy := float32(cell.Y) * ch

			drawColor := color.RGBA{20, 20, 20, 255}

			if cell.Occupied() {
				drawColor = color.RGBA{255, 255, 0, 255}
			}

			vector.StrokeLine(screen, cx, cy, cx+cw, cy, 1.0, drawColor, false)
			vector.StrokeLine(screen, cx+cw, cy, cx+cw, cy+ch, 1.0, drawColor, false)
			vector.StrokeLine(screen, cx+cw, cy+ch, cx, cy+ch, 1.0, drawColor, false)
			vector.StrokeLine(screen, cx, cy+ch, cx, cy, 1.0, drawColor, false)
		}
	}
}
