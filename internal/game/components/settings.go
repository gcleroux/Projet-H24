package components

import (
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
)

type SettingsData struct {
	Width     int
	Height    int
	Resizable bool
	CellSize  int
	ShowDebug bool
}

func GetDefaultSettingsConfig() SettingsData {
	return SettingsData{
		Width:     viper.GetInt("game.settings.window.width"),
		Height:    viper.GetInt("game.settings.window.height"),
		Resizable: viper.GetBool("game.settings.window.resizable"),
		CellSize:  viper.GetInt("game.settings.grid.cell_size"),
		ShowDebug: false,
	}
}

var Settings = donburi.NewComponentType[SettingsData]()
