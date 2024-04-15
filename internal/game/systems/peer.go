package systems

import (
	"image/color"

	nw "github.com/gcleroux/Projet-H24/internal/networking/network_client"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func DrawPeer(_ *ecs.ECS, screen *ebiten.Image) {
	peerColor := color.RGBA{0, 100, 100, 255}
	for _, peer := range nw.NetClient.Peers {
		vector.DrawFilledRect(
			screen,
			float32(peer.X),
			float32(peer.Y),
			float32(16),
			float32(24),
			peerColor,
			false,
		)
	}
}
