package scenes

import (
	"fmt"
	"image/color"

	"github.com/gcleroux/Projet-H24/assets"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/factory"
	"github.com/gcleroux/Projet-H24/internal/game/layers"
	"github.com/gcleroux/Projet-H24/internal/game/systems"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type TitleScene struct {
	ecs      *ecs.ECS
	callback func()
}

func NewTitleScene(callback func()) *TitleScene {
	s := &TitleScene{
		callback: callback,
	}
	s.configure()
	return s
}

func (s *TitleScene) Update() {
	s.ecs.Update()

	settings_entry, ok := tags.Settings.First(s.ecs.World)
	if !ok {
		return
	}
	mappings := components.KbdInput.Get(settings_entry)

	if inpututil.IsKeyJustPressed(mappings.Jump) {
		s.callback()
		return

	}
}

func (s *TitleScene) configure() {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdateSettings)

	ecs.AddRenderer(layers.LayerBackground, systems.DrawBackground)
	ecs.AddRenderer(layers.LayerDebug, systems.DrawDebug)

	s.ecs = ecs

	factory.CreateSettings(s.ecs)
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.ecs.DrawLayer(layers.LayerBackground, screen)
	s.ecs.DrawLayer(layers.LayerDebug, screen)

	settings_entry, ok := tags.Settings.First(s.ecs.World)
	if !ok {
		return
	}
	settings := components.Settings.Get(settings_entry)
	mappings := components.KbdInput.Get(settings_entry)

	text.Draw(
		screen,
		"Projet H24 - Guillaume Cléroux",
		assets.DefaultFont,
		int(settings.Width/7),
		int(settings.Height/7),
		color.White,
	)

	text.Draw(
		screen,
		fmt.Sprint(
			fmt.Sprintln("Up: "+mappings.Up.String()),
			fmt.Sprintln("Left: "+mappings.Left.String()),
			fmt.Sprintln("Down: "+mappings.Down.String()),
			fmt.Sprintln("Right: "+mappings.Right.String()),
			fmt.Sprintln("Jump: "+mappings.Jump.String()),
			fmt.Sprintln("Debug: "+mappings.ShowDebug.String()),
		),
		assets.DefaultFont,
		int(settings.Width/3),
		int(settings.Height/3),
		color.White,
	)
	text.Draw(
		screen,
		fmt.Sprintf("Press %s to start", mappings.Jump.String()),
		assets.DefaultFont,
		int(settings.Width/4),
		int(settings.Height),
		color.White,
	)
}
