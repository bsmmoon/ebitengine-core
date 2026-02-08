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

// DebugOverlay is a simple widget for displaying debug information.
type DebugOverlay struct {
	X, Y       int
	Visible    bool
	LineHeight int    // Height of each line in pixels (default: 16)
	lineCount  int    // Tracks number of lines drawn
	lines      []string // Accumulated lines to draw
}

// NewDebugOverlay creates a new DebugOverlay at the given position.
// By default, the overlay is visible.
func NewDebugOverlay(x, y int) *DebugOverlay {
	return &DebugOverlay{
		X:          x,
		Y:          y,
		Visible:    true,
		LineHeight: 16,
	}
}

// AddLine adds a line of text to the debug overlay.
// Call this multiple times to build up the debug display, then call Render().
func (d *DebugOverlay) AddLine(format string, args ...interface{}) {
	if !d.Visible {
		return
	}
	d.lines = append(d.lines, fmt.Sprintf(format, args...))
}

// Render draws all accumulated lines and clears the buffer.
// Call AddLine() multiple times, then Render() once per frame.
func (d *DebugOverlay) Render(dst *ebiten.Image) {
	if !d.Visible || len(d.lines) == 0 {
		d.lines = nil
		return
	}
	
	d.lineCount = len(d.lines)
	msg := ""
	for i, line := range d.lines {
		if i > 0 {
			msg += "\n"
		}
		msg += line
	}
	
	ebitenutil.DebugPrintAt(dst, msg, d.X, d.Y)
	d.lines = nil // Clear for next frame
}

// SetVisible sets the visibility of the debug overlay.
func (d *DebugOverlay) SetVisible(visible bool) {
	d.Visible = visible
}

// Toggle toggles the visibility of the debug overlay.
func (d *DebugOverlay) Toggle() {
	d.Visible = !d.Visible
}
