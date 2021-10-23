package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Ball struct {
	x int
	y int
	//r      int
	xSpeed int
	ySpeed int
	hitbox *resolv.Object
	logger *log.BuiltinLogger
}

func NewBall(tileSize int, logger *log.BuiltinLogger) *Ball {
	x := 32
	y := 32
	return &Ball{
		xSpeed: 1,
		ySpeed: 1,
		x:      x,
		y:      y,
		hitbox: resolv.NewObject(float64(x), float64(y), float64(tileSize), float64(tileSize)),
		logger: logger,
	}
}

func (b *Ball) Collides(l *Level) bool {
	return l.CollidesBottom(b.x, b.y)
}

func (b *Ball) Update(l *Level) error {

	if collision := b.hitbox.Check(float64(b.xSpeed), float64(b.ySpeed)); collision != nil {
		vector := collision.ContactWithObject(collision.Objects[0])
		if vector.X() != 0 {
			b.xSpeed = int(vector.X()) / 16
		}
		if vector.Y() != 0 {
			b.ySpeed = int(vector.Y()) / 16
		}
		b.logger.Infof("Collision Vector - X: %f Y: %f", vector.X(), vector.Y())

	}

	b.x += b.xSpeed
	b.y += b.ySpeed
	b.hitbox.X += float64(b.xSpeed)
	b.hitbox.Y += float64(b.ySpeed)
	b.hitbox.Update()
	return nil
}

func (b *Ball) Draw(r *ebiten.Image) {
	ebitenutil.DrawRect(r, float64(b.x), float64(b.y), 16, 16, color.White)
}
