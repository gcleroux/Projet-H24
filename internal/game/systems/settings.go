package systems

import (
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi/ecs"
)

func UpdateSettings(ecs *ecs.ECS) {
	settings := GetOrCreateSettings(ecs)

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		settings.Debug = !settings.Debug
	}
}

func GetOrCreateSettings(ecs *ecs.ECS) *components.SettingsData {
	if _, ok := components.Settings.First(ecs.World); !ok {
		ecs.World.Create(components.Settings)
	}

	ent, _ := components.Settings.First(ecs.World)
	return components.Settings.Get(ent)
}
