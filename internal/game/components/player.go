package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

// The PlayerData tracks attributes of the playable character in the current game.
// We don't have to track the position in the PlayerData itself since this is
// done in the Object component with integration with resolv
type PlayerData struct {
	SpeedX         float64
	SpeedY         float64
	OnGround       *resolv.Object
	WallSliding    *resolv.Object
	FacingRight    bool
	IgnorePlatform *resolv.Object
}

var Player = donburi.NewComponentType[PlayerData]()
