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

// IsometricProjection handles coordinate conversion for isometric rendering.
type IsometricProjection struct {
	TileSize int
}

// NewIsometricProjection creates a new IsometricProjection with the given tile size.
func NewIsometricProjection(tileSize int) *IsometricProjection {
	return &IsometricProjection{
		TileSize: tileSize,
	}
}

// CartesianToIso transforms cartesian coordinates into isometric coordinates.
// TECHNIQUE: Isometric projection - converts 2D grid position (x,y) to diamond-shaped screen position.
// Formula: ix = (x-y) * tileSize/2, iy = (x+y) * tileSize/4
func (p *IsometricProjection) CartesianToIso(x, y float64) (float64, float64) {
	ix := (x - y) * float64(p.TileSize/2)
	iy := (x + y) * float64(p.TileSize/4)
	return ix, iy
}

// IsoToCartesian transforms isometric coordinates into cartesian coordinates.
// Converts screen position back to 2D grid coordinates.
func (p *IsometricProjection) IsoToCartesian(x, y float64) (float64, float64) {
	cx := (x/float64(p.TileSize/2) + y/float64(p.TileSize/4)) / 2
	cy := (y/float64(p.TileSize/4) - (x / float64(p.TileSize/2))) / 2
	return cx, cy
}
