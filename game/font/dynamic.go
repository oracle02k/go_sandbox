package font

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// DynamicFont TrueTypeFontを使用して文字列を描画する.表示毎にテクスチャ生成〜転送を行うため動作が重い
type DynamicFont struct {
	font *ttf.Font
}

// OpenDynamicFont 使用したいTrueTypeFontを指定しDynamicFontを生成する
func OpenDynamiFont(fontPath string, size int) (*DynamicFont, error) {
	font, err := ttf.OpenFont(fontPath, size)
	if err != nil {
		return nil, err
	}

	return &DynamicFont{
		font: font,
	}, nil
}

// Close DynamicFontのリソースを解放する
func (d *DynamicFont) Close() {
	d.font.Close()
}

// Draw DynamicFontを使用して対象のRendererに文字テクスチャを描画する.表示毎にテクスチャ生成〜転送を行うため動作が重い
func (d *DynamicFont) Draw(renderer *sdl.Renderer, text string, x int32, y int32, color sdl.Color) error {
	surface, err := d.font.RenderUTF8_Blended(text, color)
	if err != nil {
		return err
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	defer texture.Destroy()

	//文字を描写したTextureのサイズを取得する
	_, _, width, height, _ := texture.Query()

	src := &sdl.Rect{X: 0, Y: 0, W: width, H: height}
	dest := &sdl.Rect{X: x, Y: y, W: width, H: height}

	renderer.Copy(texture, src, dest)

	return nil
}
