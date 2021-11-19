package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/components"
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	position *components.Position
	hitbox   *components.Hitbox
	xSpeed   float64
	ySpeed   float64
	logger   *log.BuiltinLogger
}

func NewBall(position *components.Position, logger *log.BuiltinLogger, tag string) *Ball {
	hitbox := components.NewHitbox(position, tag)
	return &Ball{
		position: position,
		xSpeed:   2.0,
		ySpeed:   2.0,
		hitbox:   hitbox,
		logger:   logger,
	}
}

func (b *Ball) Update() error {
	isXCollision, isYCollision := b.hitbox.Check(int(b.xSpeed), int(b.ySpeed))
	if isXCollision {
		b.xSpeed *= -1
	}
	if isYCollision {
		b.ySpeed *= -1
	}
	if isXCollision || isYCollision {
		b.logger.Infof("Ball - X: %d, Y: %d, isXCollision: %t, isYCollision: %t", b.position.X(), b.position.Y(), isXCollision, isYCollision)
	}

	b.position.SetX(b.position.X() + int(b.xSpeed))
	b.position.SetY(b.position.Y() + int(b.ySpeed))
	b.hitbox.UpdatePosition(int(b.xSpeed), int(b.ySpeed))

	return nil
}

func (b *Ball) Draw(r *ebiten.Image) {
	ebitenutil.DrawRect(r, float64(b.position.X()), float64(b.position.Y()), float64(b.position.Width()), float64(b.position.Height()), color.White)
}
