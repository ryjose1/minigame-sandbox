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
		xSpeed: 4,
		ySpeed: 4,
		x:      x,
		y:      y,
		hitbox: resolv.NewObject(float64(x), float64(y), float64(tileSize), float64(tileSize), "ball"),
		logger: logger,
	}
}

func (b *Ball) Update(l *Level) error {

	if collision := b.hitbox.Check(float64(b.xSpeed), float64(b.ySpeed)); collision != nil {
		for _, object := range collision.Objects {
			vector := collision.ContactWithObject(object)
			// Todo, replace with vector math
			if vector.X() == 0 {
				b.xSpeed *= -1
			}
			if vector.Y() == 0 {
				b.ySpeed *= -1
			}
			b.logger.Infof("Collision Vector - X: %f Y: %f", vector.X(), vector.Y())
		}
	}

	b.x += b.xSpeed
	b.y += b.ySpeed
	b.hitbox.X += float64(b.xSpeed)
	b.hitbox.Y += float64(b.ySpeed)
	b.hitbox.Update()
	return nil
}

func (b *Ball) Draw(r *ebiten.Image, object *resolv.Object) {
	ebitenutil.DrawRect(r, object.X, object.Y, object.W, object.H, color.White)
}
