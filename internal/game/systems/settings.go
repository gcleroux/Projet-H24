package systems

import (
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi/ecs"
)

func UpdateSettings(ecs *ecs.ECS) {
	settings_entry, _ := tags.Settings.First(ecs.World)

	settings := components.Settings.Get(settings_entry)
	mappings := components.KbdInput.Get(settings_entry)

	if inpututil.IsKeyJustPressed(mappings.ShowDebug) {
		settings.ShowDebug = !settings.ShowDebug
	}
}
