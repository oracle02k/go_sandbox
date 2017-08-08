package display

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	surface *sdl.Surface
}

func CreateTexture(filePath string) (*Texture, error) {
	surface, err := img.Load(filePath)
	if err != nil {
		return nil, err
	}

	return &Texture{
		surface: surface,
	}, nil
}

func (t *Texture) Draw(renderer *sdl.Renderer, x int32, y int32) error {

	texture, err := renderer.CreateTextureFromSurface(t.surface)
	if err != nil {
		return err
	}
	defer texture.Destroy()

	_, _, width, height, _ := texture.Query()

	src := &sdl.Rect{X: 0, Y: 0, W: width, H: height}
	dest := &sdl.Rect{X: x, Y: y, W: width, H: height}

	renderer.Copy(texture, src, dest)

	return nil
}
