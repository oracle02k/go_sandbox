package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	renderer *sdl.Renderer
}

var game = &Game{}

func initialize(renderer *sdl.Renderer) {
	game.renderer = renderer
}

func Renderer() *sdl.Renderer {
	return game.renderer
}
