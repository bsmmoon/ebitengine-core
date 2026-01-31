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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type viewport struct {
	x16 int
	y16 int
}

// Move updates the viewport position based on the background image size.
func (p *viewport) Move(width, height int) {
	maxX16 := width * 16
	maxY16 := height * 16

	// Move the viewport based on the image size.
	p.x16 += width / 32
	p.y16 += height / 32

	// Wrap the position around using modulo to create an infinite loop.
	p.x16 %= maxX16
	p.y16 %= maxY16
}

func (p *viewport) Position() (int, int) {
	return p.x16, p.y16
}

type Game struct {
	screenWidth  int
	screenHeight int
	bgImage      *ebiten.Image
	viewport     viewport
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
		bgImage:      bgImage,
	}
}

func (g *Game) Update() error {
	w, h := g.bgImage.Bounds().Dx(), g.bgImage.Bounds().Dy()
	g.viewport.Move(w, h)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x16, y16 := g.viewport.Position()
	// Convert the 1/16th pixel position to float64 for rendering.
	// We negate the position to simulate the camera moving (background moves opposite).
	offsetX, offsetY := float64(-x16)/16, float64(-y16)/16

	// Draw bgImage on the screen repeatedly.
	// A 3x3 grid ensures the screen is fully covered as the images scroll.
	const repeat = 3
	w, h := g.bgImage.Bounds().Dx(), g.bgImage.Bounds().Dy()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(g.bgImage, op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
