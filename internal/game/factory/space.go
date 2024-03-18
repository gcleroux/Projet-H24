package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/solarlune/resolv"
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)

	spaceData := resolv.NewSpace(
		viper.GetInt("window.width"),
		viper.GetInt("window.height"),
		16,
		16,
	)
	components.Space.Set(space, spaceData)

	return space
}
