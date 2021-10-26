package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

const tileSize = 16

type Level struct {
	space  *resolv.Space
	ball   *Ball
	logger *log.BuiltinLogger
	//paddle *Paddle
	//brickLayout []int
}

func NewLevel(logger *log.BuiltinLogger) *Level {
	ball := NewBall(tileSize, logger)

	return &Level{
		space:  initSpace(ball),
		ball:   ball,
		logger: logger,
	}
}

func makeBorders(width, height, tileSize float64) []*resolv.Object {
	borders := []*resolv.Object{
		// top and bottom
		resolv.NewObject(0, 0, width, tileSize, "border"),
		resolv.NewObject(0, height-tileSize, width, tileSize, "border"),
		//left and right
		resolv.NewObject(0, tileSize, tileSize, height-2*tileSize, "border"),
		resolv.NewObject(width-tileSize, tileSize, tileSize, height-2*tileSize, "border"),
	}
	return borders
}

func initSpace(ball *Ball) *resolv.Space {
	width, height := ebiten.WindowSize()
	space := resolv.NewSpace(width, height, tileSize, tileSize)

	for _, border := range makeBorders(float64(width), float64(height), float64(tileSize)) {
		space.Add(border)
	}

	space.Add(ball.hitbox)
	return space
}

func (l *Level) CollidesBottom(x, y int) bool {
	return false
}

func (l *Level) Update() error {
	l.ball.Update(l)
	return nil
}

func (l *Level) Draw(r *ebiten.Image) {
	for _, object := range l.space.Objects() {
		switch {
		case object.HasTags("border"):
			ebitenutil.DrawRect(r, object.X, object.Y, object.W, object.H, color.White)
		case object.HasTags("ball"):
			l.ball.Draw(r, object)
		}
	}

}
