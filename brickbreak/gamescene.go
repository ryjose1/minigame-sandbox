package brickbreak

import (
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	level  *Level
	logger *log.BuiltinLogger
}

func NewGameScene(logger *log.BuiltinLogger) *GameScene {
	return &GameScene{
		level:  NewLevel(logger),
		logger: logger,
	}

}

func (s *GameScene) Update() error {
	s.level.Update()
	return nil
}

func (s *GameScene) Draw(r *ebiten.Image) {
	s.level.Draw(r)
	s.logger.Infof("Ball - X:%d Y:%d %f %f", s.level.ball.x, s.level.ball.y, s.level.ball.hitbox.X, s.level.ball.hitbox.Y)
}
