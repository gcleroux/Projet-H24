package systems

import (
	"image/color"
	"math"

	"github.com/gcleroux/Projet-H24/api/v1"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/gcleroux/Projet-H24/internal/game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePlayer(ecs *ecs.ECS) {
	// Now we update the Player's movement. This is the real bread-and-butter of this example, naturally.
	playerEntry, ok := components.Player.First(ecs.World)
	if !ok {
		return
	}

	settingsEntry, ok := components.Settings.First(ecs.World)
	if !ok {
		return
	}

	player := components.Player.Get(playerEntry)
	playerObject := dresolv.GetObject(playerEntry)
	movement := components.Movement.Get(playerEntry)
	kbdMappings := components.KbdInput.Get(playerEntry)
	settings := components.Settings.Get(settingsEntry)

	player.SpeedY += movement.Gravity
	if player.WallSliding != nil && player.SpeedY > 1 {
		player.SpeedY = 1
	}

	// Horizontal movement is only possible when not wallsliding.
	if player.WallSliding == nil {
		if ebiten.IsKeyPressed(kbdMappings.Right) || ebiten.GamepadAxisValue(0, 0) > 0.1 {
			player.SpeedX += movement.Acceleration
			player.FacingRight = true
		}

		if ebiten.IsKeyPressed(kbdMappings.Left) || ebiten.GamepadAxisValue(0, 0) < -0.1 {
			player.SpeedX -= movement.Acceleration
			player.FacingRight = false
		}
	}

	// Apply friction and horizontal speed limiting.
	if player.SpeedX > movement.Friction {
		player.SpeedX -= movement.Friction
	} else if player.SpeedX < -movement.Friction {
		player.SpeedX += movement.Friction
	} else {
		player.SpeedX = 0
	}

	if player.SpeedX > movement.MaxSpeed {
		player.SpeedX = movement.MaxSpeed
	} else if player.SpeedX < -movement.MaxSpeed {
		player.SpeedX = -movement.MaxSpeed
	}

	// Check for jumping.
	if inpututil.IsKeyJustPressed(kbdMappings.Jump) || ebiten.IsGamepadButtonPressed(0, 0) ||
		ebiten.IsGamepadButtonPressed(1, 0) {
		if (ebiten.IsKeyPressed(kbdMappings.Down) || ebiten.GamepadAxisValue(0, 1) > 0.1 || ebiten.GamepadAxisValue(1, 1) > 0.1) &&
			player.OnGround != nil &&
			player.OnGround.HasTags("platform") {
			player.IgnorePlatform = player.OnGround
		} else {
			if player.OnGround != nil {
				player.SpeedY = -movement.JumpSpeed
			} else if player.WallSliding != nil {
				// WALLJUMPING
				player.SpeedY = -movement.JumpSpeed

				if player.WallSliding.Position.X > playerObject.Position.X {
					player.SpeedX = -movement.WallSpeed
				} else {
					player.SpeedX = movement.WallSpeed
				}

				player.WallSliding = nil

			}
		}
	}

	// Horizontal movement
	dx := player.SpeedX

	if check := playerObject.Check(player.SpeedX, 0, "solid"); check != nil {

		dx = check.ContactWithCell(check.Cells[0]).X
		player.SpeedX = 0

		// If you're in the air, then colliding with a wall object makes you start wall sliding.
		if player.OnGround == nil {
			player.WallSliding = check.Objects[0]
		}

	}
	playerObject.Position.X += dx

	// Vertical movement
	player.OnGround = nil
	dy := player.SpeedY
	dy = math.Max(math.Min(dy, settings.CellSize), -settings.CellSize)

	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	if check := playerObject.Check(0, checkDistance, "solid", "platform", "ramp"); check != nil {

		slide, ok := check.SlideAgainstCell(check.Cells[0], "solid")

		if dy < 0 && check.Cells[0].ContainsTags("solid") && ok &&
			math.Abs(slide.X) <= settings.CellSize/2 {
			// If we are able to slide here, we do so. No contact was made, and vertical speed (dy) is maintained upwards.
			playerObject.Position.X += slide.X
		} else {

			if ramps := check.ObjectsByTags("ramp"); len(ramps) > 0 {

				ramp := ramps[0]

				if contactSet := playerObject.Shape.Intersection(dx, settings.CellSize/2, ramp.Shape); dy >= 0 && contactSet != nil {

					// If Intersection() is successful, a ContactSet is returned. A ContactSet contains information regarding where
					// two Shapes intersect, like the individual points of contact, the center of the contacts, and the MTV, or
					// Minimum Translation Vector, to move out of contact.

					// Here, we use ContactSet.TopmostPoint() to get the top-most contact point as an indicator of where
					// we want the player's feet to be. Then we just set that position, and we're done.

					dy = contactSet.TopmostPoint().Y - playerObject.Bottom() + 0.1
					player.OnGround = ramp
					player.SpeedY = 0

				}

			}
			if platforms := check.ObjectsByTags("platform"); len(platforms) > 0 {

				platform := platforms[0]

				if platform != player.IgnorePlatform && player.SpeedY >= 0 && playerObject.Bottom() < platform.Position.Y+movement.MaxSpeed/2 {
					dy = check.ContactWithObject(platform).Y
					player.OnGround = platform
					player.SpeedY = 0
				}

			}

			if solids := check.ObjectsByTags("solid"); len(solids) > 0 && (player.OnGround == nil || player.OnGround.Position.Y >= solids[0].Position.Y) {
				dy = check.ContactWithObject(solids[0]).Y
				player.SpeedY = 0

				// We're only on the ground if we land on it (if the object's Y is greater than the player's).
				if solids[0].Position.Y > playerObject.Position.Y {
					player.OnGround = solids[0]
				}

			}

			if player.OnGround != nil {
				player.WallSliding = nil    // Player's on the ground, so no wallsliding anymore.
				player.IgnorePlatform = nil // Player's on the ground, so reset which platform is being ignored.
			}
		}
	}

	// Move the object on dy.
	playerObject.Position.Y += dy

	wallNext := 1.0
	if !player.FacingRight {
		wallNext = -1
	}

	// If the wall next to the Player runs out, stop wall sliding.
	if c := playerObject.Check(wallNext, 0, "solid"); player.WallSliding != nil && c == nil {
		player.WallSliding = nil
	}
	player.Broadcast(api.PlayerPosition{
		X: playerObject.Position.X,
		Y: playerObject.Position.Y,
	})
}

func DrawPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	e, ok := tags.Player.First(ecs.World)
	if !ok {
		return
	}
	o := dresolv.GetObject(e)
	playerColor := color.RGBA{0, 255, 60, 255}

	vector.DrawFilledRect(
		screen,
		float32(o.Position.X),
		float32(o.Position.Y),
		float32(o.Size.X),
		float32(o.Size.Y),
		playerColor,
		false,
	)
}
