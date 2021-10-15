package main

import (
	"github.com/ryjose1/minigames/brickbreak"
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	logger := log.NewBuiltinLogger()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Brick Break!")
	if err := ebiten.RunGame(brickbreak.NewGame(logger)); err != nil {
		logger.Errorf("%w", err)
	}
}
