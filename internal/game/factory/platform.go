package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlatform(ecs *ecs.ECS, object *resolv.Object) *donburi.Entry {
	platform := archetypes.Platform.Spawn(ecs)
	dresolv.SetObject(platform, object)

	return platform
}
