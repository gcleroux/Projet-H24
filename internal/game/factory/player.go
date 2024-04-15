package factory

import (
	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	nw "github.com/gcleroux/Projet-H24/internal/networking/network_client"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

	// Get start coords
	start := components.StartCoords.Get(components.StartCoords.MustFirst(ecs.World))
	settings := components.Settings.Get(components.Settings.MustFirst(ecs.World))

	obj := resolv.NewObject(
		start.X,
		start.Y,
		settings.CellSize,
		settings.CellSize*1.5,
	)
	dresolv.SetObject(player, obj)
	obj.SetShape(resolv.NewRectangle(0, 0, settings.CellSize, settings.CellSize*1.5))

	// The player will publish the position update to the NetClient
	// The NetClient will then be responsible to send the data to the server
	pub := nw.NewPublisher[api.PlayerPosition]()
	pub.AddSubscriber(nw.NetClient.Subscriber)

	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
		Publisher:   pub,
	})
	components.Movement.SetValue(player, components.GetDefaultMovementConfig())
	components.KbdInput.SetValue(player, components.GetDefaultKbdInputConfig())

	return player
}
