package factory

import (
	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	dresolv "github.com/gcleroux/Projet-H24/internal/game/resolv"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateRamp(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	ramp := archetypes.Ramp.Spawn(ecs)
	dresolv.SetObject(ramp, obj)

	// We will construct the shape using a ConvexPolygon. It's essentially an elogated triangle, but with a "floor" afterwards,
	// ensuring the Player is always able to stand regardless of which ramp they're standing on.
	rampShape := resolv.NewConvexPolygon(
		0,
		0,

		0,
		0,
		2,
		0,
		obj.Size.X-2,
		obj.Size.Y,
		obj.Size.X,
		obj.Size.Y,
		0,
		obj.Size.Y,
	)
	obj.SetShape(rampShape)

	return ramp
}
