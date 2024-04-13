package archetypes

import (
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/layers"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var (
	Platform = newArchetype(
		tags.Platform,
		components.Object,
	)
	FloatingPlatform = newArchetype(
		tags.FloatingPlatform,
		components.Object,
		components.Tween,
	)
	Player = newArchetype(
		tags.Player,
		components.Player,
		components.KbdInput,
		components.Movement,
		components.Object,
	)
	Peer = newArchetype(
		tags.Peer,
		components.Peer,
		components.Object,
	)
	Ramp = newArchetype(
		tags.Ramp,
		components.Object,
	)
	Space = newArchetype(
		components.Space,
	)
	Wall = newArchetype(
		tags.Wall,
		components.Object,
	)
	Connection = newArchetype(
		components.Connection,
	)
	Settings = newArchetype(
		tags.Settings,
		components.Settings,
		components.KbdInput,
	)
)

type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
	}
}

func (a *archetype) Spawn(ecs *ecs.ECS, cs ...donburi.IComponentType) *donburi.Entry {
	e := ecs.World.Entry(ecs.Create(
		layers.LayerActors,
		append(a.components, cs...)...,
	))
	return e
}
