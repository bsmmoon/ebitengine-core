// Copyright 2021 The Ebiten Authors
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
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Camera2D represents a 2D camera with pan and zoom capabilities.
type Camera2D struct {
	X, Y                float64
	Scale               float64
	ScaleTo             float64
	MinScale            float64
	MaxScale            float64
	ZoomSpeed           float64 // Zoom sensitivity (default: 7.0)
	PanSpeed            float64 // Pan speed (default: 7.0)
	MousePanSensitivity float64 // Mouse pan sensitivity (default: 100.0)
	mousePanX           int
	mousePanY           int
	interpolation       float64
}

// NewCamera2D creates a new Camera2D with default settings.
func NewCamera2D() *Camera2D {
	return &Camera2D{
		Scale:               1.0,
		ScaleTo:             1.0,
		MinScale:            0.01,
		MaxScale:            100.0,
		ZoomSpeed:           7.0,
		PanSpeed:            7.0,
		MousePanSensitivity: 100.0,
		mousePanX:           math.MinInt32,
		mousePanY:           math.MinInt32,
		interpolation:       10.0,
	}
}

// HandleZoom processes zoom input and updates target scale.
func (c *Camera2D) HandleZoom(scrollY float64) {
	c.ScaleTo += scrollY * (c.ScaleTo / c.ZoomSpeed)
	if c.ScaleTo < c.MinScale {
		c.ScaleTo = c.MinScale
	} else if c.ScaleTo > c.MaxScale {
		c.ScaleTo = c.MaxScale
	}
}

// HandleKeyboardPan processes keyboard pan input.
func (c *Camera2D) HandleKeyboardPan(dx, dy float64) {
	pan := c.PanSpeed / c.Scale
	c.X += dx * pan
	c.Y += dy * pan
}

// HandleMousePan processes mouse drag panning.
// Returns true if currently panning.
func (c *Camera2D) HandleMousePan() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if c.mousePanX == math.MinInt32 && c.mousePanY == math.MinInt32 {
			c.mousePanX, c.mousePanY = ebiten.CursorPosition()
		} else {
			x, y := ebiten.CursorPosition()
			pan := c.PanSpeed / c.Scale
			dx, dy := float64(c.mousePanX-x)*(pan/c.MousePanSensitivity), float64(c.mousePanY-y)*(pan/c.MousePanSensitivity)
			c.X, c.Y = c.X-dx, c.Y+dy
		}
		return true
	} else if c.mousePanX != math.MinInt32 || c.mousePanY != math.MinInt32 {
		c.mousePanX, c.mousePanY = math.MinInt32, math.MinInt32
	}
	return false
}

// Update smoothly interpolates scale toward target.
// TECHNIQUE: Smooth zoom transition - creates gradual zoom effect instead of instant jumps.
func (c *Camera2D) Update() {
	if c.ScaleTo > c.Scale {
		c.Scale += (c.ScaleTo - c.Scale) / c.interpolation
	} else if c.ScaleTo < c.Scale {
		c.Scale -= (c.Scale - c.ScaleTo) / c.interpolation
	}
}

// Clamp restricts camera position to world bounds.
func (c *Camera2D) Clamp(minX, minY, maxX, maxY float64) {
	if c.X < minX {
		c.X = minX
	} else if c.X > maxX {
		c.X = maxX
	}
	if c.Y < minY {
		c.Y = minY
	} else if c.Y > maxY {
		c.Y = maxY
	}
}
