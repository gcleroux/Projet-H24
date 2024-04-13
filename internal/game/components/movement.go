package components

import (
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
)

type MovementData struct {
	MaxSpeed, JumpSpeed, WallSpeed, Friction, Acceleration, Gravity float64
}

func GetDefaultMovementConfig() MovementData {
	return MovementData{
		MaxSpeed:     viper.GetFloat64("game.settings.movement.max_speed"),
		JumpSpeed:    viper.GetFloat64("game.settings.movement.jump_speed"),
		WallSpeed:    viper.GetFloat64("game.settings.movement.wall_speed"),
		Friction:     viper.GetFloat64("game.settings.movement.friction"),
		Acceleration: viper.GetFloat64("game.settings.movement.acceleration"),
		Gravity:      viper.GetFloat64("game.settings.movement.gravity"),
	}
}

var Movement = donburi.NewComponentType[MovementData]()
