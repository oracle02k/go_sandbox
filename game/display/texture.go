package display

import (
	"github.com/oracle02k/go_sandbox/game"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	surface *sdl.Surface
	texture *sdl.Texture
	width   int32
	height  int32
}

func CreateTexture(filePath string) (*Texture, error) {
	surface, err := img.Load(filePath)
	if err != nil {
		return nil, err
	}

	texture, err := game.Renderer().CreateTextureFromSurface(surface)
	if err != nil {
		surface.Free()
		return nil, err
	}

	_, _, width, height, err := texture.Query()
	if err != nil {
		texture.Destroy()
		surface.Free()
		return nil, err
	}

	return &Texture{
		surface: surface,
		texture: texture,
		width:   width,
		height:  height,
	}, nil
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
	t.surface.Free()
}

func (t *Texture) Draw(x int32, y int32) {

	src := &sdl.Rect{X: 0, Y: 0, W: t.width, H: t.height}
	dest := &sdl.Rect{X: x, Y: y, W: t.width, H: t.height}

	game.Renderer().Copy(t.texture, src, dest)
}
