package brickbreak

import (
	"image/color"

	"github.com/ryjose1/minigames/components"
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	position *components.Position
	objects  *LevelObjects
	space    *components.HitSpace
	logger   *log.BuiltinLogger

	//brickLayout []int
}

// NewLevel creates a new brickbreak level
func NewLevel(position *components.Position, logger *log.BuiltinLogger) *Level {
	objects := NewLevelObjects(logger, position)
	borders := makeLevelBorders(position, 16)
	space := components.NewSpace(position)

	space.AddHitboxes(append([]*components.Hitbox{objects.paddle.hitbox, objects.ball.hitbox}, borders...))

	return &Level{
		position: position,
		objects:  objects,
		space:    space,
		logger:   logger,
	}
}

type LevelObjects struct {
	ball   *Ball
	paddle *Paddle
}

func NewLevelObjects(logger *log.BuiltinLogger, levelPosition *components.Position) *LevelObjects {
	ballX := levelPosition.X() + int(float64(levelPosition.Width())*0.1)
	ballY := levelPosition.Y() + int(float64(levelPosition.Height())*0.1)
	ballPosition := components.NewPosition(ballX, ballY, components.TILESIZE, components.TILESIZE)
	ball := NewBall(ballPosition, logger, "ball")

	paddleX := levelPosition.X() + int(float64(levelPosition.Width())*0.5)
	paddleY := levelPosition.Y() + int(float64(levelPosition.Height())*0.8)
	paddlePosition := components.NewPosition(paddleX, paddleY, components.TILESIZE*4, components.TILESIZE)
	paddle := NewPaddle(paddlePosition, logger, "paddle")

	return &LevelObjects{
		ball:   ball,
		paddle: paddle,
	}
}

// makeLevelBorders creates the border objects, which serve as the "walls" in the level
func makeLevelBorders(levelPosition *components.Position, thickness int) []*components.Hitbox {
	width := levelPosition.Width()
	height := levelPosition.Height()
	x := levelPosition.X()
	y := levelPosition.Y()

	borders := []*components.Hitbox{
		components.NewHitbox(components.NewPosition(x, y, width, thickness), "border"),
		components.NewHitbox(components.NewPosition(x, y+height-thickness, width, thickness), "border"),
		components.NewHitbox(components.NewPosition(x, y+thickness, thickness, height-2*thickness), "border"),
		components.NewHitbox(components.NewPosition(x+width-thickness, y+thickness, thickness, height-2*thickness), "border"),
	}
	return borders
}

// Update progresses the objects in the level by one tick
func (l *Level) Update() error {
	l.objects.ball.Update()
	l.objects.paddle.Update("border")
	return nil
}

// Draw creates visualizations for each of the objects in the level
func (l *Level) Draw(r *ebiten.Image) {

	l.objects.paddle.Draw(r)
	l.objects.ball.Draw(r)
	for _, object := range l.space.Objects() {
		switch {
		case object.HasTags("border"):
			ebitenutil.DrawRect(r, object.X, object.Y, object.W, object.H, color.White)
		}
	}

}
