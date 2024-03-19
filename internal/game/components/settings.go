package components

import "github.com/yohamta/donburi"

type SettingsData struct {
	Debug bool
}

var Settings = donburi.NewComponentType[SettingsData]()
