package game_test

import (
	"math"
	"testing"

	"github.com/gcleroux/Projet-H24/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestPlayer_Update(t *testing.T) {
	type fields struct {
		Image       *ebiten.Image
		Position    game.Position
		Speed       float64
		Velocity    float64
		UpperBoundX float64
		UpperBoundY float64
	}
	type args struct {
		inputs []ebiten.Key
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantX, wantY float64
	}{
		{
			name: "Update with no inputs",
			fields: fields{
				Position: game.Position{100, 100},
				Speed:    5.0,
			},
			args: args{
				inputs: []ebiten.Key{},
			},
			wantErr: false,
			wantX:   100.0,
			wantY:   100.0,
		},
		{
			name: "Update KeyW",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyW},
			},
			wantErr: false,
			wantX:   100.0,
			wantY:   95.0,
		},
		{
			name: "Update KeyA",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyA},
			},
			wantErr: false,
			wantX:   95.0,
			wantY:   100.0,
		},
		{
			name: "Update KeyS",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyS},
			},
			wantErr: false,
			wantX:   100.0,
			wantY:   105.0,
		},
		{
			name: "Update KeyD",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyD},
			},
			wantErr: false,
			wantX:   105.0,
			wantY:   100.0,
		},
		{
			name: "Two keys cancel each other",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyA, ebiten.KeyD},
			},
			wantErr: false,
			wantX:   100.0,
			wantY:   100.0,
		},
		{
			name: "Diagonal movement doesn't make you go faster",
			fields: fields{
				Position:    game.Position{100, 100},
				Speed:       5.0,
				UpperBoundX: 200,
				UpperBoundY: 200,
			},
			args: args{
				inputs: []ebiten.Key{ebiten.KeyW, ebiten.KeyD},
			},
			wantErr: false,
			wantX:   100 + 5.0/math.Sqrt2,
			wantY:   100 - 5.0/math.Sqrt2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &game.Player{
				Image:       tt.fields.Image,
				Position:    tt.fields.Position,
				Speed:       tt.fields.Speed,
				Velocity:    tt.fields.Velocity,
				UpperBoundX: tt.fields.UpperBoundX,
				UpperBoundY: tt.fields.UpperBoundY,
			}
			if err := p.Update(tt.args.inputs); (err != nil) != tt.wantErr {
				t.Errorf("Player.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check the actual position against the expected position
			if p.Position.X != tt.wantX || p.Position.Y != tt.wantY {
				t.Errorf(
					"Player position after update = (%v, %v), want (%v, %v)",
					p.Position.X,
					p.Position.Y,
					tt.wantX,
					tt.wantY,
				)
			}
		})
	}
}
