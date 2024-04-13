package factory

import (
	"log"

	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)
	settings_entry, ok := tags.Settings.First(ecs.World)
	if !ok {
		log.Fatal("Unable to find Settings components")
	}
	settings := components.Settings.Get(settings_entry)

	spaceData := resolv.NewSpace(
		settings.Width,
		settings.Height,
		settings.CellSize,
		settings.CellSize,
	)
	components.Space.Set(space, spaceData)

	return space
}
