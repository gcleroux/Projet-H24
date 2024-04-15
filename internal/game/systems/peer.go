package systems

import (
	"image/color"

	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	nw "github.com/gcleroux/Projet-H24/internal/networking/network_client"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func DrawPeer(ecs *ecs.ECS, screen *ebiten.Image) {
	settings_entry, ok := tags.Settings.First(ecs.World)
	if !ok {
		return
	}
	settings := components.Settings.Get(settings_entry)

	peerColor := color.RGBA{0, 100, 100, 255}
	for _, peer := range nw.NetClient.Peers {
		vector.DrawFilledRect(
			screen,
			float32(peer.X),
			float32(peer.Y),
			float32(settings.CellSize),
			float32(settings.CellSize*1.5),
			peerColor,
			false,
		)
	}
}
