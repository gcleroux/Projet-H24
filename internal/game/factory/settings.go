package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSettings(ecs *ecs.ECS) *donburi.Entry {
	settings := archetypes.Settings.Spawn(ecs)

	components.Settings.SetValue(settings, components.GetDefaultSettingsConfig())
	components.KbdInput.SetValue(settings, components.GetDefaultKbdInputConfig())

	return settings
}
