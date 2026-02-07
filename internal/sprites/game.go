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

package sprites

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math"
	"math/rand/v2"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	MinSprites = 0
	MaxSprites = 50000
)

type Game struct {
	debugui      debugui.DebugUI
	sprites      Sprites
	op           ebiten.DrawImageOptions
	inited       bool
	screenWidth  int
	screenHeight int
	ebitenImage  *ebiten.Image
}

func NewGame(cfg GameConfig) *Game {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	origEbitenImage := ebiten.NewImageFromImage(img)

	s := origEbitenImage.Bounds().Size()
	ebitenImage := ebiten.NewImage(s.X, s.Y)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.5)
	ebitenImage.DrawImage(origEbitenImage, op)

	return &Game{
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
		ebitenImage:  ebitenImage,
	}
}

// init initializes the sprites with random positions, velocities, and angles.
func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	g.sprites.sprites = make([]*Sprite, MaxSprites)
	g.sprites.num = 500
	for i := range g.sprites.sprites {
		w, h := g.ebitenImage.Bounds().Dx(), g.ebitenImage.Bounds().Dy()
		x, y := rand.IntN(g.screenWidth-w), rand.IntN(g.screenHeight-h)
		vx, vy := 2*rand.IntN(2)-1, 2*rand.IntN(2)-1
		a := rand.IntN(maxAngle)
		g.sprites.sprites[i] = &Sprite{
			imageWidth:  w,
			imageHeight: h,
			x:           x,
			y:           y,
			vx:          vx,
			vy:          vy,
			angle:       a,
		}
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	// Update the debug UI (allows changing the number of sprites).
	if _, err := g.debugui.Update(func(ctx *debugui.Context) error {
		ctx.Window("Sprites", image.Rect(10, 10, 210, 110), func(layout debugui.ContainerLayout) {
			ctx.Text(fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
			ctx.Text(fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
			ctx.Slider(&g.sprites.num, 0, 50000, 100)
		})
		return nil
	}); err != nil {
		return err
	}

	g.sprites.Update(g.screenWidth, g.screenHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.ebitenImage.Bounds().Dx(), g.ebitenImage.Bounds().Dy()
	for i := 0; i < g.sprites.num; i++ {
		s := g.sprites.sprites[i]
		// Reset the geometry matrix for the current sprite.
		g.op.GeoM.Reset()
		// Move origin to the center of the image to rotate around the center.
		g.op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		g.op.GeoM.Rotate(2 * math.Pi * float64(s.angle) / maxAngle)
		// Move origin back and translate to the sprite's position.
		g.op.GeoM.Translate(float64(w)/2, float64(h)/2)
		g.op.GeoM.Translate(float64(s.x), float64(s.y))
		screen.DrawImage(g.ebitenImage, &g.op)
	}

	g.debugui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
