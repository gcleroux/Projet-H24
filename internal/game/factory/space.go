package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)

	settings_entry := tags.Settings.MustFirst(ecs.World)
	settings := components.Settings.Get(settings_entry)

	spaceData := resolv.NewSpace(
		int(settings.Width),
		int(settings.Height),
		int(settings.CellSize),
		int(settings.CellSize),
	)
	components.Space.Set(space, spaceData)

	return space
}
