package game

import "github.com/veandco/go-sdl2/sdl"

// Scene ゲームシーンのインターフェイス
type Scene interface {
	Progress()
	Draw(renderer *sdl.Renderer)
}
