package components

import "github.com/yohamta/donburi"

type SettingsData struct {
	Width     int
	Height    int
	Resizable bool
	Debug     bool
}

var Settings = donburi.NewComponentType[SettingsData]()
