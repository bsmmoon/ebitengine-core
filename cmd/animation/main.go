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

package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	// frameOX and frameOY are the origin (top-left corner) of the animation sequence in the spritesheet.
	frameOX = 0
	frameOY = 32
	// frameWidth and frameHeight are the dimensions of a single frame.
	frameWidth  = 32
	frameHeight = 32
	// frameCount is the total number of frames in the animation sequence.
	frameCount = 8
)

var (
	// runnerImage stores the loaded spritesheet image.
	runnerImage *ebiten.Image
)

type Game struct {
	// count is used to track the elapsed time (in ticks) for animation timing.
	count int
}

func (g *Game) Update() error {
	// Update the tick counter.
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// Move the image's origin to its center so we can position it by its center.
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	// Move the image to the center of the screen.
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	// Calculate the current frame index based on the game tick count.
	// Dividing by 5 slows down the animation (updates every 5 ticks).
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY

	// Extract the sub-image corresponding to the current frame from the spritesheet.
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
