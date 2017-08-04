package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Run 指定されたシーンを起点としてゲームを開始する
func Run(title string, root Scene) error {
	if sdl.Init(sdl.INIT_EVERYTHING) != nil {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	if ttf.Init() != nil {
		panic(sdl.GetError())
	}
	defer ttf.Quit()

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	loop(root, renderer)

	return nil
}

func loop(root Scene, renderer *sdl.Renderer) {
	var nextFrame = sdl.GetTicks()
	var wait = 1000.0 / 60

	for pollEvent() == 0 {
		/* 1秒間に60回Updateされるようにする */
		if sdl.GetTicks() >= nextFrame {
			//update()
			root.Progress()
			/* 時間がまだあるときはDrawする */
			if sdl.GetTicks() < nextFrame+uint32(wait) {
				renderer.Clear()
				renderer.SetDrawColor(128, 0, 0, 255)
				renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: 800, H: 600})
				root.Draw(renderer)
				renderer.Present()
			}
			nextFrame += uint32(wait)
			sdl.Delay(0)
		}
	}
}

func pollEvent() int {
	result := 0
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			result = 1
		case *sdl.KeyUpEvent:
			switch t.Keysym.Sym {
			case sdl.K_ESCAPE:
				result = 1
			}
		}
	}

	return result
}
