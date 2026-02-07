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

const (
	maxAngle = 256
)

// Sprite represents a moving image with position, velocity, and rotation.
type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          int
	vy          int
	angle       int
}

// Update moves the sprite and bounces it off the screen edges.
func (s *Sprite) Update(screenWidth, screenHeight int) {
	s.x += s.vx
	s.y += s.vy
	// Check horizontal bounds and bounce if necessary.
	if s.x < 0 {
		s.x = -s.x
		s.vx = -s.vx
	} else if mx := screenWidth - s.imageWidth; mx <= s.x {
		s.x = 2*mx - s.x
		s.vx = -s.vx
	}
	// Check vertical bounds and bounce if necessary.
	if s.y < 0 {
		s.y = -s.y
		s.vy = -s.vy
	} else if my := screenHeight - s.imageHeight; my <= s.y {
		s.y = 2*my - s.y
		s.vy = -s.vy
	}
	// Update rotation angle.
	s.angle++
	if s.angle == maxAngle {
		s.angle = 0
	}
}

// Sprites manages a collection of sprites.
type Sprites struct {
	sprites []*Sprite
	num     int
}

func (s *Sprites) Update(screenWidth, screenHeight int) {
	for i := 0; i < s.num; i++ {
		s.sprites[i].Update(screenWidth, screenHeight)
	}
}
