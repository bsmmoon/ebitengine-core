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

package shared

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents an animated sprite from a spritesheet.
type Sprite struct {
	image       *ebiten.Image
	frameWidth  int
	frameHeight int
	startX      int
	startY      int
	frameCount  int
	speed       int // Ticks per frame
	tick        int
}

// NewSprite creates a new animated sprite.
func NewSprite(img *ebiten.Image, frameWidth, frameHeight, startX, startY, frameCount, speed int) *Sprite {
	return &Sprite{
		image:       img,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		startX:      startX,
		startY:      startY,
		frameCount:  frameCount,
		speed:       speed,
	}
}

// Update advances the animation timer.
func (s *Sprite) Update() {
	s.tick++
}

// Draw renders the current frame of the sprite at the given position.
func (s *Sprite) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	// Calculate current frame index
	i := 0
	if s.speed > 0 {
		i = (s.tick / s.speed) % s.frameCount
	}
	sx := s.startX + i*s.frameWidth
	sy := s.startY

	// Extract and draw the sub-image
	subImg := s.image.SubImage(image.Rect(sx, sy, sx+s.frameWidth, sy+s.frameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)
}

// Size returns the width and height of a single frame.
func (s *Sprite) Size() (int, int) {
	return s.frameWidth, s.frameHeight
}
