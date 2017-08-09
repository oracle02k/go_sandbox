package scene

import (
	"github.com/oracle02k/go_sandbox/game"
	"github.com/oracle02k/go_sandbox/game/display"
)

type sceneTexture struct {
	texture *display.Texture
}

func NewSceneTexture() game.Scene {
	return &sceneTexture{}
}

// Initialize シーン初期化
func (s *sceneTexture) Initialize() error {

	texture, _ := display.CreateTexture("il_32.png")
	s.texture = texture

	return nil
}

// Progress シーン初期化
func (s *sceneTexture) Progress() {
}

// Renderer シーン初期化
func (s *sceneTexture) Draw() {
	s.texture.Draw(10, 10)
}
