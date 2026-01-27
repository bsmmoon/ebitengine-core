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

// TileMap is a reusable component for rendering tile-based maps with multiple layers.
type TileMap struct {
	Spritesheet *Spritesheet
	Layers      [][]int // Multiple layers of tile indices (0 = transparent/skip)
	Width       int     // Map width in tiles
	X, Y        int     // Position offset in pixels
}

// NewTileMap creates a new TileMap with the given spritesheet and map width.
func NewTileMap(spritesheet *Spritesheet, width int) *TileMap {
	return &TileMap{
		Spritesheet: spritesheet,
		Layers:      make([][]int, 0),
		Width:       width,
		X:           0,
		Y:           0,
	}
}

// AddLayer adds a new layer of tiles to the map.
// Each layer is a slice of tile indices in row-major order.
// Tile index 0 is treated as transparent/skip in overlay layers.
func (t *TileMap) AddLayer(tiles []int) {
	t.Layers = append(t.Layers, tiles)
}

// SetPosition sets the position offset of the tilemap in pixels.
func (t *TileMap) SetPosition(x, y int) {
	t.X = x
	t.Y = y
}

// Draw renders all layers of the tilemap to the destination image.
// Layers are drawn in order (first layer at bottom, last layer on top).
func (t *TileMap) Draw(dst *ebiten.Image) {
	tileSize := t.Spritesheet.TileSize

	for layerIdx, layer := range t.Layers {
		for i, tileIndex := range layer {
			// Skip tile index 0 for overlay layers (index > 0)
			// This allows the background layer to show through
			if layerIdx > 0 && tileIndex == 0 {
				continue
			}

			// Calculate destination position
			tileX := i % t.Width
			tileY := i / t.Width
			dstX := float64(t.X + tileX*tileSize)
			dstY := float64(t.Y + tileY*tileSize)

			// Get the tile image from the spritesheet
			tileImg := t.Spritesheet.TileAt(tileIndex)

			// Draw the tile
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(dstX, dstY)
			dst.DrawImage(tileImg, op)
		}
	}
}

// Height returns the height of the tilemap in tiles.
func (t *TileMap) Height() int {
	if len(t.Layers) == 0 || len(t.Layers[0]) == 0 {
		return 0
	}
	return len(t.Layers[0]) / t.Width
}

// PixelWidth returns the width of the tilemap in pixels.
func (t *TileMap) PixelWidth() int {
	return t.Width * t.Spritesheet.TileSize
}

// PixelHeight returns the height of the tilemap in pixels.
func (t *TileMap) PixelHeight() int {
	return t.Height() * t.Spritesheet.TileSize
}
