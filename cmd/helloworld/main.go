package main

import (
	"log"

	"github.com/bsmmoon/ebitengine-core/internal/helloworld"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(helloworld.NewGame()); err != nil {
		log.Fatal(err)
	}
}
