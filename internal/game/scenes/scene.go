package scenes

import (
	"sync"

	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/events"
	"github.com/gcleroux/Projet-H24/internal/game/factory"
	"github.com/gcleroux/Projet-H24/internal/game/layers"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/systems"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type PlatformerScene struct {
	ecs  *ecs.ECS
	once sync.Once
}

func (ps *PlatformerScene) Update() {
	ps.once.Do(ps.configure)
	ps.ecs.Update()
}

func (ps *PlatformerScene) Draw(screen *ebiten.Image) {
	ps.ecs.DrawLayer(layers.LayerBackground, screen)
	ps.ecs.DrawLayer(layers.LayerActors, screen)
	ps.ecs.DrawLayer(layers.LayerDebug, screen)
}

func (ps *PlatformerScene) configure() {
	ecs := ecs.NewECS(donburi.NewWorld())

	// Update systems
	ecs.AddSystem(systems.UpdatePlayer)
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
	events.PeerUpdateEvent.Subscribe(ecs.World, events.PeerUpdateHandler)
	events.PlayerUpdateEvent.Subscribe(ecs.World, events.PlayerUpdateHandler)

	ps.ecs = ecs

	factory.CreateSettings(ps.ecs)
	factory.CreateConnection(ps.ecs)

	settings := components.Settings.Get(tags.Settings.MustFirst(ps.ecs.World))
	gw, gh := float64(settings.Width), float64(settings.Height)

	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	space := factory.CreateSpace(ps.ecs)

	dresolv.Add(space,
		// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells,
		// as it all is in this platformer example.

		// TODO: Use something else to create the maps instead of hardcoding values here
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-16, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, gw, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, gh-24, gw, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(160, gh-56, 160, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(320, 64, 32, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(64, 128, 16, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, 64, 128, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, gh-88, 128, 16, "solid")),

		// Create the Player. NewPlayer adds it to the world's Space.
		factory.CreatePlayer(ps.ecs),

		// Create the Peers
		// TODO: Don't preallocate peers on the scene, or maybe wait for full lobby before loading scene
		factory.CreatePeer(ps.ecs),
		factory.CreatePeer(ps.ecs),
		factory.CreatePeer(ps.ecs),

		// Non-moving floating Platforms.
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+64, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+128, 48, 8, "platform")),
		factory.CreatePlatform(ps.ecs, resolv.NewObject(352, 64+192, 48, 8, "platform")),

		// // Create the floating platform.
		// factory.CreateFloatingPlatform(ps.ecs, resolv.NewObject(128, gh-32, 128, 8, "platform")),
		// // A ramp, which is unique as it has a non-rectangular shape. For this, we will specify a different shape for collision testing.
		// factory.CreateRamp(ps.ecs, resolv.NewObject(320, gh-56, 64, 32, "ramp")),
	)
}
