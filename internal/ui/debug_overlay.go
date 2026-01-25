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

package ui

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
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
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
