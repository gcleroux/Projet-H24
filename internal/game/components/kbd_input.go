package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
)

// KbdInputData are the mappings from string to ebiten.Key
type KbdInputData struct {
	Left, Right, Up, Down, Jump, ShowDebug ebiten.Key
}

func GetDefaultKbdInputConfig() KbdInputData {
	k := KbdInputData{}
	if err := k.Left.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.left"))); err != nil {
		log.Println(err)
	}
	if err := k.Right.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.right"))); err != nil {
		log.Println(err)
	}
	if err := k.Up.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.up"))); err != nil {
		log.Println(err)
	}
	if err := k.Down.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.down"))); err != nil {
		log.Println(err)
	}
	if err := k.Jump.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.jump"))); err != nil {
		log.Println(err)
	}
	if err := k.ShowDebug.UnmarshalText([]byte(viper.GetString("game.keyboard_mappings.show_debug"))); err != nil {
		log.Println(err)
	}
	return k
}

var KbdInput = donburi.NewComponentType[KbdInputData]()
