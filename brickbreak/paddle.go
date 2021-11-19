package brickbreak

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ryjose1/minigames/components"
	"github.com/ryjose1/minigames/log"
)

type Paddle struct {
	position *components.Position
	hitbox   *components.Hitbox
	xSpeed   int
	ySpeed   int

	logger *log.BuiltinLogger
}

func NewPaddle(position *components.Position, logger *log.BuiltinLogger, tag string) *Paddle {
	hitbox := components.NewHitbox(position, tag)
	return &Paddle{
		position: position,
		xSpeed:   2,
		ySpeed:   0,
		hitbox:   hitbox,
		logger:   logger,
	}
}

func (p *Paddle) Update(wallTag string) error {
	isXCollision, isYCollision := p.hitbox.Check(p.xSpeed, p.ySpeed, wallTag)

	if isXCollision {
		p.xSpeed *= -1
	}

	p.position.SetX(p.position.X() + int(p.xSpeed))
	p.hitbox.UpdatePosition(p.xSpeed, 0)
	if isXCollision || isYCollision {
		p.logger.Infof("Paddle - X: %d, Y: %d, isXCollision: %t, isYCollision: %t", p.position.X(), p.position.Y(), isXCollision, isYCollision)
	}
	return nil
}

func (p *Paddle) Draw(r *ebiten.Image) {
	ebitenutil.DrawRect(r, float64(p.position.X()), float64(p.position.Y()), float64(p.position.Width()), float64(p.position.Height()), color.White)
}
