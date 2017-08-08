package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/oracle02k/go_sandbox/game"
	"github.com/oracle02k/go_sandbox/game/scene"
)

func main() {
	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	*/
	if err := game.Run("sandbox", scene.NewSceneTexture()); err != nil {
		log.Fatal(err)
	}
}
