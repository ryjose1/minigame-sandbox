package brickbreak

import (
	"github.com/ryjose1/minigames/components"
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	level  *Level
	logger *log.BuiltinLogger
}

func NewGameScene(logger *log.BuiltinLogger) *GameScene {
	width, height := ebiten.WindowSize()

	levelPosition := components.NewPosition(0, 0, width/4*3, height)

	return &GameScene{
		level:  NewLevel(levelPosition, logger),
		logger: logger,
	}

}

func (s *GameScene) Update() error {
	s.level.Update()
	return nil
}

func (s *GameScene) Draw(r *ebiten.Image) {
	s.level.Draw(r)
}
