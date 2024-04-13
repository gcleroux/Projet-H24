package scenes

import (
	"github.com/gcleroux/Projet-H24/internal/game/scenes/levels"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update()
	Draw(*ebiten.Image)
}

type SceneManager struct {
	current Scene
}

func NewSceneManager() *SceneManager {
	s := &SceneManager{}
	s.SwitchToTitleScene()
	return s
}

func (s *SceneManager) Update() {
	s.current.Update()
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	s.current.Draw(screen)
}

func (s *SceneManager) SwitchToTitleScene() {
	s.current = NewTitleScene(s.SwitchToGameScene)
	// s.current = &TitleScene{
	// 	callback: s.SwitchToGameScene,
	// }
}

func (s *SceneManager) SwitchToGameScene() {
	s.current = levels.NewLevel_00_Scene()
}

// func (s *SceneManager) NextScene() {
// 	cu
// }
