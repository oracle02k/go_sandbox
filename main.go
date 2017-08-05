package main

import (
	"log"

	"github.com/oracle02k/go_sandbox/game"
	"github.com/oracle02k/go_sandbox/game/scene"
)

func main() {
	if err := game.Run("sandbox", &scene.Text{}); err != nil {
		log.Fatal(err)
	}
}
