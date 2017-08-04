package main

import (
	"github.com/oracle02k/go_sandbox/game"
	"github.com/oracle02k/go_sandbox/game/scene"
)

func main() {
	game.Run("sandbox", &scene.Text{})
}
