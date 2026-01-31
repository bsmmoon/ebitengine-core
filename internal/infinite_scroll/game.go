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

package infinite_scroll

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/bsmmoon/ebitengine-core/internal/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Game struct {
	screenWidth  int
	screenHeight int
	background   *shared.ScrollingBackground
}

func NewGame(cfg GameConfig) *Game {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Tile_png))
	if err != nil {
		log.Fatal(err)
	}
	bgImage := ebiten.NewImageFromImage(img)

	return &Game{
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
		background:   shared.NewScrollingBackground(bgImage),
	}
}

func (g *Game) Update() error {
	w, _ := g.background.Size()
	g.background.Update(w/32, 0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
