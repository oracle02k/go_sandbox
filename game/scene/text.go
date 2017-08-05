package scene

import (
	"github.com/oracle02k/go_sandbox/game/font"
	"github.com/veandco/go-sdl2/sdl"
)

// Text テキストシーン
type Text struct {
	font *font.DynamicFont
	x    int32
}

// Initialize シーン初期化
func (t *Text) Initialize() error {
	font, err := font.OpenDynamiFont("Arial.ttf", 20)
	if err != nil {
		return err
	}

	t.font = font
	return nil
}

// Progress シーン更新
func (t *Text) Progress() {
	t.x++
}

// Draw シーン描画
func (t *Text) Draw(renderer *sdl.Renderer) {
	t.font.Draw(renderer, "hoge Hoge", t.x, 10, sdl.Color{R: 255, G: 255, B: 255, A: 255})
}
