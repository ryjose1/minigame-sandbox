package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Wall struct {
	position *components.Position
	hitbox   *components.Hitbox
}

func NewWall(position *components.Position, offset *components.Position, tag string) *Wall {
	hitbox := components.NewHitbox(components.NewPosition(position.X()-offset.X(), position.Y()-offset.Y(), position.Width(), position.Height()), tag)
	return &Wall{
		position: position,
		hitbox:   hitbox,
	}
}

func (w *Wall) Draw(r *ebiten.Image) {
	ebitenutil.DrawRect(r, float64(w.position.X()), float64(w.position.Y()), float64(w.position.Width()), float64(w.position.Height()), color.White)
}
