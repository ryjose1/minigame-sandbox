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

func initSpace(ball *Ball) *resolv.Space {
	width, height := ebiten.WindowSize()
	space := resolv.NewSpace(width, height, tileSize, tileSize)
	space.Add(
		resolv.NewObject(0, 0, 640, 16),
		resolv.NewObject(0, 480-16, 640, 16),
		resolv.NewObject(0, 16, 16, 480-32),
		resolv.NewObject(640-16, 16, 16, 480-32),
	)

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
	ebitenutil.DrawRect(r, 0, 0, 640, 16, color.White)
	ebitenutil.DrawRect(r, 0, 480-16, 640, 16, color.White)
	ebitenutil.DrawRect(r, 0, 16, 16, 480-32, color.White)
	ebitenutil.DrawRect(r, 640-16, 16, 16, 480-32, color.White)
	l.ball.Draw(r)
}
