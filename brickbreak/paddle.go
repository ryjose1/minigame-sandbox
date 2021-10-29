package brickbreak

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ryjose1/minigames/log"
	"github.com/solarlune/resolv"
)

type Paddle struct {
	x      float64
	y      float64
	xSpeed float64
	ySpeed float64
	hitbox *resolv.Object

	logger *log.BuiltinLogger
}

func NewPaddle(tileSize float64, logger *log.BuiltinLogger, tag string) *Paddle {
	x := 320.0
	y := 400.0
	width := tileSize * 4
	height := tileSize
	return &Paddle{
		x:      x,
		y:      y,
		xSpeed: 2.0,
		hitbox: resolv.NewObject(x, y, width, height, tag),
		logger: logger,
	}
}

func (p *Paddle) Update(wallTag string) error {

	if collision := p.hitbox.Check(p.xSpeed, p.ySpeed, wallTag); collision != nil {
		for _, object := range collision.Objects {
			vector := collision.ContactWithObject(object)
			// Todo, replace with vector math
			if vector.X() == 0 {
				p.xSpeed *= -1
			}
		}
	}

	p.x += p.xSpeed
	p.hitbox.X += p.xSpeed
	p.hitbox.Update()

	return nil
}

func (p *Paddle) Draw(r *ebiten.Image, object *resolv.Object) {
	ebitenutil.DrawRect(r, object.X, object.Y, object.W, object.H, color.White)
}
