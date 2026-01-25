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

package tiles

import (
	"image"
	"log"

	"github.com/bsmmoon/ebitengine-core/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	tileSize  = 16
	mapWidth  = 15 // Map width in tiles
	mapHeight = 15 // Map height in tiles
)

// Game implements ebiten.Game interface for the tiles demo.
type Game struct {
	screenWidth    int
	screenHeight   int
	tileMap        *ui.TileMap
	debugOverlay   *ui.DebugOverlay
	interactions   *InteractionManager
	lastClickedObj string
}

// NewGame creates a new tiles game with the given configuration.
func NewGame(cfg GameConfig) *Game {
	// Load the tiles spritesheet
	tilesImage, err := ui.LoadImageFromBytes(images.Tiles_png)
	if err != nil {
		log.Fatal(err)
	}

	// Create spritesheet from the loaded image
	spritesheet := ui.NewSpritesheet(tilesImage, tileSize)

	// Create tilemap with the spritesheet
	tileMap := ui.NewTileMap(spritesheet, mapWidth)

	// Add background layer (grass with decorations)
	tileMap.AddLayer(parseCSVLayer(backgroundLayerData))

	// Add foreground layer (house and path)
	tileMap.AddLayer(parseCSVLayer(foregroundLayerData))

	// Create debug overlay at top-left
	debugOverlay := ui.NewDebugOverlay(0, 0)

	// Create interaction manager
	interactions := NewInteractionManager(tileSize)
	
	g := &Game{
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
		tileMap:      tileMap,
		debugOverlay: debugOverlay,
		interactions: interactions,
	}
	
	// Set interaction callback
	interactions.SetOnInteraction(func(name string, tileX, tileY int) {
		log.Printf("%s clicked at tile: (%d, %d)", name, tileX, tileY)
		g.lastClickedObj = name + " clicked!"
	})
	interactions.AddObject(InteractiveObject{
		Name:   "House",
		Bounds: image.Rect(5, 1, 11, 6),
	})

	return g
}

// Update implements ebiten.Game interface.
func (g *Game) Update() error {
	g.interactions.Update()
	return nil
}

// Draw implements ebiten.Game interface.
func (g *Game) Draw(screen *ebiten.Image) {
	g.tileMap.Draw(screen)
	g.debugOverlay.Draw(screen)
	
	// Show last clicked object message
	if g.lastClickedObj != "" {
		g.debugOverlay.DrawMessage(screen, g.lastClickedObj, 0, 20)
	}
}

// Layout implements ebiten.Game interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
