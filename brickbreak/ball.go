package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Ball struct {
	x float64
	y float64
	//r      int
	xSpeed float64
	ySpeed float64
	hitbox *resolv.Object
	logger *log.BuiltinLogger
}

func NewBall(tileSize float64, logger *log.BuiltinLogger, tag string) *Ball {
	x := 32.0
	y := 32.0
	return &Ball{
		x:      x,
		y:      y,
		xSpeed: 4.0,
		ySpeed: 4.0,
		hitbox: resolv.NewObject(x, y, tileSize, tileSize, tag),
		logger: logger,
	}
}

func (b *Ball) Update() error {

	if collision := b.hitbox.Check(b.xSpeed, b.ySpeed); collision != nil {
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
	b.hitbox.X += b.xSpeed
	b.hitbox.Y += b.ySpeed
	b.hitbox.Update()
	return nil
}

func (b *Ball) Draw(r *ebiten.Image, object *resolv.Object) {
	ebitenutil.DrawRect(r, object.X, object.Y, object.W, object.H, color.White)
}
