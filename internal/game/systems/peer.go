package systems

import (
	"image/color"

	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawPeer(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Peer.Each(ecs.World, func(e *donburi.Entry) {
		peer := components.Peer.Get(e)

		// Only draw the peer when in game
		if peer.Present {
			o := dresolv.GetObject(e)
			peerColor := color.RGBA{0, 255, 60, 255}
			vector.DrawFilledRect(
				screen,
				float32(o.Position.X),
				float32(o.Position.Y),
				float32(o.Size.X),
				float32(o.Size.Y),
				peerColor,
				false,
			)
		}
	})
}
