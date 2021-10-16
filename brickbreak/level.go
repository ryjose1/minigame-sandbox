package brickbreak

import "github.com/hajimehoshi/ebiten/v2"

type Level struct {
	windowHeight int
	windowWidth  int
}

func NewLevel() *Level {
	width, height := ebiten.WindowSize()
	return &Level{
		windowWidth:  width,
		windowHeight: height,
	}
}

func (l *Level) CollidesBottom(x, y int) bool {
	return y >= l.windowHeight
}
