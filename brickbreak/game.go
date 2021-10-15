package brickbreak

import (
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *SceneManager
	Logger       *log.BuiltinLogger
}

func NewGame(logger *log.BuiltinLogger) *Game {
	return &Game{
		Logger:       logger,
		sceneManager: NewSceneManager(NewGameScene(logger)),
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if err := g.sceneManager.Update(); err != nil {
		return err
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.sceneManager.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
