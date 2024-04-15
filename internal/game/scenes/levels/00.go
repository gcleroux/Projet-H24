package levels

import (
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/factory"
	"github.com/gcleroux/Projet-H24/internal/game/layers"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Level_00_Scene struct {
	ecs *ecs.ECS
}

func NewLevel_00_Scene() *Level_00_Scene {
	s := &Level_00_Scene{}
	s.configure()
	return s
}

func (s *Level_00_Scene) Update() {
	s.ecs.Update()
}

func (s *Level_00_Scene) Draw(screen *ebiten.Image) {
	s.ecs.DrawLayer(layers.LayerBackground, screen)
	s.ecs.DrawLayer(layers.LayerActors, screen)
	s.ecs.DrawLayer(layers.LayerDebug, screen)
}

func (s *Level_00_Scene) configure() {
	ecs := ecs.NewECS(donburi.NewWorld())

	// Update systems
	ecs.AddSystem(systems.UpdatePlayer)
	// ecs.AddSystem(systems.UpdatePeer)
	// ecs.AddSystem(systems.UpdateConnection)
	ecs.AddSystem(systems.UpdateObjects)
	ecs.AddSystem(systems.UpdateSettings)

	// Rendering systems
	ecs.AddRenderer(layers.LayerActors, systems.DrawWall)
	ecs.AddRenderer(layers.LayerActors, systems.DrawPlatform)
	ecs.AddRenderer(layers.LayerActors, systems.DrawPlayer)
	ecs.AddRenderer(layers.LayerActors, systems.DrawPeer)
	ecs.AddRenderer(layers.LayerBackground, systems.DrawBackground)
	ecs.AddRenderer(layers.LayerDebug, systems.DrawDebug)

	// Events handling
	// events.PeerUpdateEvent.Subscribe(ecs.World, events.PeerUpdateHandler)
	// events.PlayerUpdateEvent.Subscribe(ecs.World, events.PlayerUpdateHandler)

	s.ecs = ecs
	// factory.CreateConnection(s.ecs)
	settings := components.Settings.Get(factory.CreateSettings(s.ecs))
	gw, gh := float64(settings.Width), float64(settings.Height)

	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	space := factory.CreateSpace(s.ecs)

	dresolv.Add(space,
		// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells,
		// as it all is in this platformer example.

		// TODO: Use something else to create the maps instead of hardcoding values here
		factory.CreateWall(s.ecs, resolv.NewObject(0, 0, 16, gh, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(gw-16, 0, 16, gh, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(0, 0, gw, 16, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(0, gh-24, gw, 32, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(160, gh-56, 160, 32, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(320, 64, 32, 160, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(64, 128, 16, 160, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(gw-128, 64, 128, 16, "solid")),
		factory.CreateWall(s.ecs, resolv.NewObject(gw-128, gh-88, 128, 16, "solid")),

		// Create the Player. NewPlayer adds it to the world's Space.
		factory.CreatePlayer(s.ecs),

		// Non-moving floating Platforms.
		factory.CreatePlatform(s.ecs, resolv.NewObject(352, 64, 48, 8, "platform")),
		factory.CreatePlatform(s.ecs, resolv.NewObject(352, 64+64, 48, 8, "platform")),
		factory.CreatePlatform(s.ecs, resolv.NewObject(352, 64+128, 48, 8, "platform")),
		factory.CreatePlatform(s.ecs, resolv.NewObject(352, 64+192, 48, 8, "platform")),
	)
}
