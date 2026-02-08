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

package isometric

import (
	"fmt"

	"github.com/bsmmoon/ebitengine-core/internal/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game is an isometric demo game.
type Game struct {
	w, h         int
	currentLevel *Level
	camera       *shared.Camera2D
	projection   *shared.IsometricProjection
	debugOverlay *shared.DebugOverlay
	offscreen    *ebiten.Image
}

// NewGame returns a new isometric demo Game.
func NewGame(cfg GameConfig) (*Game, error) {
	l, err := NewLevel()
	if err != nil {
		return nil, fmt.Errorf("failed to create new level: %s", err)
	}

	g := &Game{
		w:            cfg.ScreenWidth,
		h:            cfg.ScreenHeight,
		currentLevel: l,
		camera:       shared.NewCamera2D(),
		projection:   shared.NewIsometricProjection(l.tileSize),
		debugOverlay: shared.NewDebugOverlay(10, 10),
	}
	return g, nil
}

// Update reads current user input and updates the Game state.
func (g *Game) Update() error {
	// Handle zoom input.
	var scrollY float64
	if ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		scrollY = -0.25
	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		scrollY = .25
	} else {
		_, scrollY = ebiten.Wheel()
		if scrollY < -1 {
			scrollY = -1
		} else if scrollY > 1 {
			scrollY = 1
		}
	}
	g.camera.HandleZoom(scrollY)

	// Update camera (smooth zoom interpolation happens here).
	g.camera.Update()

	// Handle keyboard panning.
	var dx, dy float64
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		dx = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		dy = 1
	}
	g.camera.HandleKeyboardPan(dx, dy)

	// Handle mouse panning.
	g.camera.HandleMousePan()

	// Clamp camera position.
	worldWidth := float64(g.currentLevel.w * g.currentLevel.tileSize / 2)
	worldHeight := float64(g.currentLevel.h * g.currentLevel.tileSize / 2)
	g.camera.Clamp(-worldWidth, -worldHeight, worldWidth, 0)

	// Randomize level.
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		l, err := NewLevel()
		if err != nil {
			return fmt.Errorf("failed to create new level: %s", err)
		}
		g.currentLevel = l
		g.projection = shared.NewIsometricProjection(l.tileSize)
	}

	return nil
}

// Draw draws the Game on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Render level.
	g.renderLevel(screen)

	// Build debug overlay
	g.debugOverlay.AddLine("KEYS WASD EC R")
	g.debugOverlay.AddLine("SCA  %0.2f", g.camera.Scale)
	g.debugOverlay.AddLine("POS  %0.0f,%0.0f", g.camera.X, g.camera.Y)
	g.debugOverlay.AddLine("TPS: %0.2f", ebiten.ActualTPS())
	g.debugOverlay.AddLine("FPS: %0.2f", ebiten.ActualFPS())
	g.debugOverlay.Render(screen)
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.w, g.h = outsideWidth, outsideHeight
	return g.w, g.h
}

// renderLevel draws the current Level on the screen.
func (g *Game) renderLevel(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	padding := float64(g.currentLevel.tileSize) * g.camera.Scale
	cx, cy := float64(g.w/2), float64(g.h/2)

	scaleLater := g.camera.Scale > 1
	target := screen
	scale := g.camera.Scale

	// TECHNIQUE: Anti-bleeding - When zooming in (>1x), render at 1x scale to offscreen buffer
	// first, then scale the final result. This prevents pixel bleeding between tiles.
	// Without this, scaled tiles can show thin lines between them due to floating-point rounding.
	if scaleLater {
		if g.offscreen != nil {
			if g.offscreen.Bounds().Size() != screen.Bounds().Size() {
				g.offscreen.Deallocate()
				g.offscreen = nil
			}
		}
		if g.offscreen == nil {
			s := screen.Bounds().Size()
			g.offscreen = ebiten.NewImage(s.X, s.Y)
		}
		target = g.offscreen
		target.Clear()
		scale = 1
	}

	// TECHNIQUE: Depth sorting - Iterate Y first (outer loop), then X (inner loop).
	// This ensures tiles render back-to-front using painter's algorithm,
	// so tiles in front properly overlap tiles behind them in isometric view.
	for y := 0; y < g.currentLevel.h; y++ {
		for x := 0; x < g.currentLevel.w; x++ {
			xi, yi := g.projection.CartesianToIso(float64(x), float64(y))

			// TECHNIQUE: Frustum culling - Skip tiles outside the visible screen area.
			// This optimization renders only ~5% of tiles (those actually visible),
			// dramatically improving performance for large levels.
			drawX, drawY := ((xi-g.camera.X)*g.camera.Scale)+cx, ((yi+g.camera.Y)*g.camera.Scale)+cy
			if drawX+padding < 0 || drawY+padding < 0 || drawX > float64(g.w) || drawY > float64(g.h) {
				continue
			}

			t := g.currentLevel.tiles[y][x]
			if t == nil {
				continue // No tile at this position.
			}

			op.GeoM.Reset()
			// Move to current isometric position.
			op.GeoM.Translate(xi, yi)
			// Translate camera position.
			op.GeoM.Translate(-g.camera.X, g.camera.Y)
			// Zoom.
			op.GeoM.Scale(scale, scale)
			// Center.
			op.GeoM.Translate(cx, cy)

			t.Draw(target, op)
		}
	}

	if scaleLater {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-cx, -cy)
		op.GeoM.Scale(float64(g.camera.Scale), float64(g.camera.Scale))
		op.GeoM.Translate(cx, cy)
		screen.DrawImage(target, op)
	}
}
