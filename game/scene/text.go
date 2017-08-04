package scene

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Text struct{}

func (t *Text) Progress()                   {}
func (t *Text) Draw(renderer *sdl.Renderer) {}
