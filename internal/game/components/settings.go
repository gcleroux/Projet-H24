package components

import (
	"github.com/spf13/viper"
	"github.com/yohamta/donburi"
)

type SettingsData struct {
	Width              float64
	Height             float64
	Resizable          bool
	CellSize           float64
	LatencyUpdateFrame int
	Ticker             int
	ClientLatency      int64
	ServerLatency      int64
	TotalLatency       int64
	ShowDebug          bool
}

func GetDefaultSettingsConfig() SettingsData {
	return SettingsData{
		Width:              viper.GetFloat64("game.settings.window.width"),
		Height:             viper.GetFloat64("game.settings.window.height"),
		Resizable:          viper.GetBool("game.settings.window.resizable"),
		CellSize:           viper.GetFloat64("game.settings.grid.cell_size"),
		LatencyUpdateFrame: viper.GetInt("game.settings.window.latency_update_frame"),
		Ticker:             0,
		ClientLatency:      0,
		ServerLatency:      0,
		TotalLatency:       0,
		ShowDebug:          false,
	}
}

var Settings = donburi.NewComponentType[SettingsData]()
