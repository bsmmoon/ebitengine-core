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
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

// LoadImageFromBytes decodes an image from a byte slice and returns an ebiten.Image.
func LoadImageFromBytes(data []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

// Spritesheet helps extract sub-images (tiles) from a tile atlas.
type Spritesheet struct {
	Image    *ebiten.Image
	TileSize int
	// Calculated values
	tilesPerRow int
}

// NewSpritesheet creates a new Spritesheet from an ebiten.Image with the given tile size.
func NewSpritesheet(img *ebiten.Image, tileSize int) *Spritesheet {
	bounds := img.Bounds()
	tilesPerRow := bounds.Dx() / tileSize
	return &Spritesheet{
		Image:       img,
		TileSize:    tileSize,
		tilesPerRow: tilesPerRow,
	}
}

// TileAt returns the tile at the given linear index (row-major order).
// Index 0 is the top-left tile, index 1 is the next tile to the right, etc.
func (s *Spritesheet) TileAt(index int) *ebiten.Image {
	x := index % s.tilesPerRow
	y := index / s.tilesPerRow
	return s.TileAtXY(x, y)
}

// TileAtXY returns the tile at the given grid coordinates.
// (0, 0) is the top-left tile.
func (s *Spritesheet) TileAtXY(x, y int) *ebiten.Image {
	sx := x * s.TileSize
	sy := y * s.TileSize
	return s.Image.SubImage(image.Rect(sx, sy, sx+s.TileSize, sy+s.TileSize)).(*ebiten.Image)
}

// TilesPerRow returns the number of tiles per row in the spritesheet.
func (s *Spritesheet) TilesPerRow() int {
	return s.tilesPerRow
}
