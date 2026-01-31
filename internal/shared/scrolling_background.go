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
	"github.com/hajimehoshi/ebiten/v2"
)

// ScrollingBackground represents a background image that scrolls infinitely.
type ScrollingBackground struct {
	image *ebiten.Image
	x16   int
	y16   int
}

// NewScrollingBackground creates a new ScrollingBackground.
func NewScrollingBackground(img *ebiten.Image) *ScrollingBackground {
	return &ScrollingBackground{
		image: img,
	}
}

// Update moves the scroll position by (dx, dy) in 1/16th pixels.
func (s *ScrollingBackground) Update(dx16, dy16 int) {
	w, h := s.image.Bounds().Dx(), s.image.Bounds().Dy()
	maxX16 := w * 16
	maxY16 := h * 16

	s.x16 += dx16
	s.y16 += dy16

	s.x16 %= maxX16
	s.y16 %= maxY16
}

// Draw draws the background image repeatedly to cover the screen.
func (s *ScrollingBackground) Draw(screen *ebiten.Image) {
	x16, y16 := s.x16, s.y16
	offsetX, offsetY := float64(-x16)/16, float64(-y16)/16

	const repeat = 3
	w, h := s.image.Bounds().Dx(), s.image.Bounds().Dy()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(s.image, op)
		}
	}
}

// Size returns the size of the background image.
func (s *ScrollingBackground) Size() (int, int) {
	return s.image.Bounds().Dx(), s.image.Bounds().Dy()
}
