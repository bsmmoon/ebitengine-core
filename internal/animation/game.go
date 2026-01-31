// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package animation

import (
	"log"

	"github.com/bsmmoon/ebitengine-core/internal/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Game struct {
	screenWidth  int
	screenHeight int
	runnerSprite *shared.Sprite
}

func NewGame(cfg GameConfig) *Game {
	// Use shared utility to load image
	img, err := shared.LoadImageFromBytes(images.Runner_png)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new sprite: 32x32 frames, starting at (0, 32), 8 frames, speed 5
	sprite := shared.NewSprite(img, 32, 32, 0, 32, 8, 5)

	return &Game{
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
		runnerSprite: sprite,
	}
}

func (g *Game) Update() error {
	g.runnerSprite.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.runnerSprite.Size()
	// Center the sprite on the screen
	x := float64(g.screenWidth)/2 - float64(w)/2
	y := float64(g.screenHeight)/2 - float64(h)/2
	g.runnerSprite.Draw(screen, x, y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
