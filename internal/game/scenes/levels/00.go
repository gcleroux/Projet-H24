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

	s.ecs = ecs

	settings := components.Settings.Get(factory.CreateSettings(s.ecs))
	gw, gh := settings.Width, settings.Height

	// Create the startCoords for the scene
	start := ecs.World.Entry(ecs.Create(layers.LayerActors, components.StartCoords))
	components.StartCoords.SetValue(start, components.StartCoordsData{
		X: settings.CellSize * 2,
		Y: gh - (settings.CellSize * 2),
	})

	// Define the world's Space. Here, a Space is essentially a grid, made up of settings.CellSizexsettings.CellSize cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	space := factory.CreateSpace(s.ecs)

	dresolv.Add(
		space,
		// Construct the solid level geometry.
		// Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells

		// TODO: Use something else to create the maps instead of hardcoding values here
		factory.CreateWall(s.ecs, resolv.NewObject(0, 0, settings.CellSize, gh, "solid")),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(gw-settings.CellSize, 0, settings.CellSize, gh, "solid"),
		),
		factory.CreateWall(s.ecs, resolv.NewObject(0, 0, gw, settings.CellSize, "solid")),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(0, gh-(settings.CellSize*1.5), gw, settings.CellSize*2, "solid"),
		),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(
				settings.CellSize*10,
				gh-(settings.CellSize*3.5),
				settings.CellSize*10,
				settings.CellSize*2,
				"solid",
			),
		),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(
				settings.CellSize*20,
				settings.CellSize*4,
				settings.CellSize*2,
				settings.CellSize*10,
				"solid",
			),
		),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(
				settings.CellSize*4,
				settings.CellSize*8,
				settings.CellSize,
				settings.CellSize*10,
				"solid",
			),
		),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(
				gw-(settings.CellSize*8),
				settings.CellSize*4,
				settings.CellSize*8,
				settings.CellSize,
				"solid",
			),
		),
		factory.CreateWall(
			s.ecs,
			resolv.NewObject(
				gw-(settings.CellSize*8),
				gh-(settings.CellSize*5.5),
				settings.CellSize*8,
				settings.CellSize,
				"solid",
			),
		),

		// Create the Player. NewPlayer adds it to the world's Space.
		factory.CreatePlayer(s.ecs),

		// Non-moving floating Platforms.
		factory.CreatePlatform(
			s.ecs,
			resolv.NewObject(
				gh-(settings.CellSize/2),
				settings.CellSize*4,
				settings.CellSize*3,
				settings.CellSize/2,
				"platform",
			),
		),
		factory.CreatePlatform(
			s.ecs,
			resolv.NewObject(
				gh-(settings.CellSize/2),
				settings.CellSize*8,
				settings.CellSize*3,
				settings.CellSize/2,
				"platform",
			),
		),
		factory.CreatePlatform(
			s.ecs,
			resolv.NewObject(
				gh-(settings.CellSize/2),
				settings.CellSize*12,
				settings.CellSize*3,
				settings.CellSize/2,
				"platform",
			),
		),
		factory.CreatePlatform(
			s.ecs,
			resolv.NewObject(
				gh-(settings.CellSize/2),
				settings.CellSize*16,
				settings.CellSize*3,
				settings.CellSize/2,
				"platform",
			),
		),
	)
}
