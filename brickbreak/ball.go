package brickbreak

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	x int
	y int
	//r      int
	xSpeed int
	ySpeed int
}

func NewBall() *Ball {
	return &Ball{
		xSpeed: 1,
		ySpeed: 1,
	}
}

func (b *Ball) Collides(l *Level) bool {
	return l.CollidesBottom(b.x, b.y)
}

func (b *Ball) Update(l *Level) error {
	if b.Collides(l) {
		b.ySpeed *= -1
	}

	b.x += b.xSpeed
	b.y += b.ySpeed
	return nil
}

func (b *Ball) Draw(r *ebiten.Image) {
	ebitenutil.DrawRect(r, float64(b.x), float64(b.y), 10, 10, color.White)
}
