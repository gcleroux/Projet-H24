package components

import (
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
)

type SettingsData struct {
	Width              int
	Height             int
	Resizable          bool
	CellSize           int
	LatencyUpdateFrame int
	Ticker             int
	ClientLatency      int64
	ServerLatency      int64
	TotalLatency       int64
	ShowDebug          bool
}

func GetDefaultSettingsConfig() SettingsData {
	return SettingsData{
		Width:              viper.GetInt("game.settings.window.width"),
		Height:             viper.GetInt("game.settings.window.height"),
		Resizable:          viper.GetBool("game.settings.window.resizable"),
		LatencyUpdateFrame: viper.GetInt("game.settings.window.latency_update_frame"),
		Ticker:             0,
		ClientLatency:      0,
		ServerLatency:      0,
		TotalLatency:       0,
		CellSize:           viper.GetInt("game.settings.grid.cell_size"),
		ShowDebug:          false,
	}
}

var Settings = donburi.NewComponentType[SettingsData]()
