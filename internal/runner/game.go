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

package runner

import (
	"log"

	"github.com/bsmmoon/ebitengine-core/internal/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Game struct {
	screenWidth    int
	screenHeight   int
	standingSprite *shared.Sprite
	runningSprite  *shared.Sprite
	activeSprite   *shared.Sprite
}

func NewGame(cfg GameConfig) *Game {
	// Use shared utility to load image
	img, err := shared.LoadImageFromBytes(images.Runner_png)
	if err != nil {
		log.Fatal(err)
	}

	standingSprite := shared.NewSprite(img, 32, 32, 0, 0, 5, 5) // standing
	runningSprite := shared.NewSprite(img, 32, 32, 0, 32, 8, 5) // running

	return &Game{
		screenWidth:    cfg.ScreenWidth,
		screenHeight:   cfg.ScreenHeight,
		standingSprite: standingSprite,
		runningSprite:  runningSprite,
		activeSprite:   standingSprite,
	}
}

func (g *Game) Update() error {
	flipH := g.activeSprite.FlipH
	moving := false

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		flipH = true
		moving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		flipH = false
		moving = true
	}

	if moving {
		g.activeSprite = g.runningSprite
	} else {
		g.activeSprite = g.standingSprite
	}
	g.activeSprite.FlipH = flipH
	g.activeSprite.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.activeSprite.Size()
	// Center the sprite on the screen
	x := float64(g.screenWidth)/2 - float64(w)/2
	y := float64(g.screenHeight)/2 - float64(h)/2
	g.activeSprite.Draw(screen, x, y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
