package brickbreak

import (
	"github.com/ryjose1/minigames/components"
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
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
	space := components.NewSpace(position)

	wallHitboxes := []*components.Hitbox{}
	for _, wall := range objects.walls {
		wallHitboxes = append(wallHitboxes, wall.hitbox)
	}

	space.AddHitboxes(append([]*components.Hitbox{objects.paddle.hitbox, objects.ball.hitbox}, wallHitboxes...))

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
	walls  []*Wall
}

func NewLevelObjects(logger *log.BuiltinLogger, levelPosition *components.Position) *LevelObjects {
	// TODO: Fix collisions for imperfect overlaps
	//ballX := levelPosition.X() + int(float64(levelPosition.Width())*0.1)
	//ballY := levelPosition.Y() + int(float64(levelPosition.Height())*0.1)
	ballX := levelPosition.X() + int(float64(levelPosition.Width())*0.25)
	ballY := levelPosition.Y() + int(float64(levelPosition.Height())*0.25)
	ballPosition := components.NewPosition(ballX, ballY, components.TILESIZE, components.TILESIZE)
	ball := NewBall(ballPosition, levelPosition, logger, "ball")

	paddleX := levelPosition.X() + int(float64(levelPosition.Width())*0.5)
	paddleY := levelPosition.Y() + int(float64(levelPosition.Height())*0.8)
	paddlePosition := components.NewPosition(paddleX, paddleY, components.TILESIZE*4, components.TILESIZE)
	paddle := NewPaddle(paddlePosition, levelPosition, logger, "paddle")

	walls := makeLevelWalls(levelPosition, 16)

	return &LevelObjects{
		ball:   ball,
		paddle: paddle,
		walls:  walls,
	}
}

// makeLevelWalls creates the border objects, which serve as the "walls" in the level
func makeLevelWalls(levelPosition *components.Position, thickness int) []*Wall {
	width := levelPosition.Width()
	height := levelPosition.Height()
	x := levelPosition.X()
	y := levelPosition.Y()

	borders := []*Wall{
		NewWall(components.NewPosition(x, y, width, thickness), levelPosition, "border"),
		NewWall(components.NewPosition(x, y+height-thickness, width, thickness), levelPosition, "border"),
		NewWall(components.NewPosition(x, y+thickness, thickness, height-2*thickness), levelPosition, "border"),
		NewWall(components.NewPosition(x+width-thickness, y+thickness, thickness, height-2*thickness), levelPosition, "border"),
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
	for _, wall := range l.objects.walls {
		wall.Draw(r)
	}
}
