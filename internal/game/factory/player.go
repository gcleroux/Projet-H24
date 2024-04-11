package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

	// TODO: We could add a resolv tag like "character" or "solid"
	// TODO: The spawn coords should be assigned by the server
	obj := resolv.NewObject(32, 128, 16, 24)
	dresolv.SetObject(player, obj)
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 16, 24))

	return player
}
