package brickbreak

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
	next    Scene
}

func NewSceneManager(current Scene) *SceneManager {
	return &SceneManager{
		current: current,
	}
}

func (s *SceneManager) Update() error {
	return s.current.Update()
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

func (s *SceneManager) GoTo(scene Scene) {
	s.current = scene
}
