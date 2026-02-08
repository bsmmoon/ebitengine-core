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
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DebugOverlay is a simple widget for displaying TPS/FPS debug information.
type DebugOverlay struct {
	X, Y    int
	Visible bool
}

// NewDebugOverlay creates a new DebugOverlay at the given position.
// By default, the overlay is visible.
func NewDebugOverlay(x, y int) *DebugOverlay {
	return &DebugOverlay{
		X:       x,
		Y:       y,
		Visible: true,
	}
}

// Draw renders the debug overlay showing TPS information.
func (d *DebugOverlay) Draw(dst *ebiten.Image) {
	if !d.Visible {
		return
	}
	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrintAt(dst, msg, d.X, d.Y)
}

// SetVisible sets the visibility of the debug overlay.
func (d *DebugOverlay) SetVisible(visible bool) {
	d.Visible = visible
}

// Toggle toggles the visibility of the debug overlay.
func (d *DebugOverlay) Toggle() {
	d.Visible = !d.Visible
}

// DrawWithText renders the debug overlay with custom text prepended.
func (d *DebugOverlay) DrawWithText(dst *ebiten.Image, text string) {
	if !d.Visible {
		return
	}
	msg := fmt.Sprintf("%s\nTPS: %0.2f\nFPS: %0.2f", text, ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrintAt(dst, msg, d.X, d.Y)
}

// DrawMessage renders a custom message at the specified offset from the overlay position.
func (d *DebugOverlay) DrawMessage(dst *ebiten.Image, msg string, offsetX, offsetY int) {
	if !d.Visible {
		return
	}
	ebitenutil.DebugPrintAt(dst, msg, d.X+offsetX, d.Y+offsetY)
}
