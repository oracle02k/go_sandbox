package scene

import (
	"fmt"

	"github.com/oracle02k/go_sandbox/game"
	"github.com/oracle02k/go_sandbox/game/font"
	"github.com/veandco/go-sdl2/sdl"
)

type sceneDeviceInfo struct {
	font *font.DynamicFont
}

// NewSceneDeviceInfo シーン構造体の作成
func NewSceneDeviceInfo() game.Scene {
	return &sceneDeviceInfo{}
}

// Initialize シーン初期化
func (s *sceneDeviceInfo) Initialize() error {
	font, err := font.OpenDynamiFont("Arial.ttf", 10)
	if err != nil {
		return err
	}

	s.font = font
	return nil
}

// Progress シーン更新
func (s *sceneDeviceInfo) Progress() {
}

// Draw シーン描画
func (s *sceneDeviceInfo) Draw(renderer *sdl.Renderer) {

	info := sdl.RendererInfo{}
	renderer.GetRendererInfo(&info)

	names := map[int32]string{
		sdl.PIXELFORMAT_UNKNOWN:     "PIXELFORMAT_UNKNOWN",
		sdl.PIXELFORMAT_INDEX1LSB:   "PIXELFORMAT_INDEX1LSB",
		sdl.PIXELFORMAT_INDEX1MSB:   "PIXELFORMAT_INDEX1MSB",
		sdl.PIXELFORMAT_INDEX4LSB:   "PIXELFORMAT_INDEX4LSB",
		sdl.PIXELFORMAT_INDEX4MSB:   "PIXELFORMAT_INDEX4M",
		sdl.PIXELFORMAT_INDEX8:      "PIXELFORMAT_INDEX8",
		sdl.PIXELFORMAT_RGB332:      "PIXELFORMAT_RGB332",
		sdl.PIXELFORMAT_RGB444:      "PIXELFORMAT_RGB444",
		sdl.PIXELFORMAT_RGB555:      "PIXELFORMAT_RGB555",
		sdl.PIXELFORMAT_BGR555:      "PIXELFORMAT_BGR555",
		sdl.PIXELFORMAT_ARGB4444:    "PIXELFORMAT_ARGB4444",
		sdl.PIXELFORMAT_RGBA4444:    "PIXELFORMAT_RGBA4444",
		sdl.PIXELFORMAT_ABGR4444:    "PIXELFORMAT_ABGR4444",
		sdl.PIXELFORMAT_BGRA4444:    "PIXELFORMAT_BGRA4444",
		sdl.PIXELFORMAT_ARGB1555:    "PIXELFORMAT_ARGB1555",
		sdl.PIXELFORMAT_RGBA5551:    "PIXELFORMAT_RGBA5551",
		sdl.PIXELFORMAT_ABGR1555:    "PIXELFORMAT_ABGR1555",
		sdl.PIXELFORMAT_BGRA5551:    "PIXELFORMAT_BGRA5551",
		sdl.PIXELFORMAT_RGB565:      "PIXELFORMAT_RGB565",
		sdl.PIXELFORMAT_BGR565:      "PIXELFORMAT_BGR565",
		sdl.PIXELFORMAT_RGB24:       "PIXELFORMAT_RGB24",
		sdl.PIXELFORMAT_BGR24:       "PIXELFORMAT_BGR24",
		sdl.PIXELFORMAT_RGB888:      "PIXELFORMAT_RGB888",
		sdl.PIXELFORMAT_RGBX8888:    "PIXELFORMAT_RGBX8888",
		sdl.PIXELFORMAT_BGR888:      "PIXELFORMAT_BGR888",
		sdl.PIXELFORMAT_BGRX8888:    "PIXELFORMAT_BGRX8888",
		sdl.PIXELFORMAT_ARGB8888:    "PIXELFORMAT_ARGB8888",
		sdl.PIXELFORMAT_RGBA8888:    "PIXELFORMAT_RGBA8888",
		sdl.PIXELFORMAT_ABGR8888:    "PIXELFORMAT_ABGR8888",
		sdl.PIXELFORMAT_BGRA8888:    "PIXELFORMAT_BGRA8888",
		sdl.PIXELFORMAT_ARGB2101010: "PIXELFORMAT_ARGB2101010",
		sdl.PIXELFORMAT_YV12:        "PIXELFORMAT_YV12",
		sdl.PIXELFORMAT_IYUV:        "PIXELFORMAT_IYUV",
		sdl.PIXELFORMAT_YUY2:        "PIXELFORMAT_YUY2",
		sdl.PIXELFORMAT_UYVY:        "PIXELFORMAT_UYVY",
		sdl.PIXELFORMAT_YVYU:        "PIXELFORMAT_YVYU",
	}

	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	s.font.Draw(renderer, fmt.Sprintf("device name: %v", info.Name), 10, 10, color)
	s.font.Draw(renderer, fmt.Sprintf("flags: %v", info.Flags), 10, 22, color)
	s.font.Draw(renderer, fmt.Sprintf("max texture width: %v", info.MaxTextureWidth), 10, 32, color)
	s.font.Draw(renderer, fmt.Sprintf("max texture height: %v", info.MaxTextureHeight), 10, 42, color)
	s.font.Draw(renderer, fmt.Sprintf("num texture formats: %v", info.NumTextureFormats), 10, 52, color)

	for i := 0; i < len(info.TextureFormats); i++ {
		if info.TextureFormats[i] != 0 {
			y := int32(62 + i*10)
			value := info.TextureFormats[i]
			name := names[value]
			s.font.Draw(renderer, fmt.Sprintf("texture format: %v(%v)", name, value), 10, y, color)
		}
	}
}
