package shared

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InteractiveObject represents a clickable rectangular area in the game.
// The Bounds are specified in grid coordinates (e.g., tile coordinates for tile-based games,
// or pixel coordinates for pixel-based games, depending on the tileSize used in InteractionManager).
type InteractiveObject struct {
	Name    string
	Bounds  image.Rectangle // Bounds in grid coordinates (see InteractionManager.tileSize)
	OnClick func()          // Optional callback when this object is clicked
}

// InteractionManager handles click detection for rectangular interactive objects.
// It converts mouse pixel coordinates to grid coordinates using tileSize.
//
// Usage examples:
//   - Tile-based games: tileSize=16, bounds in tiles (e.g., Rect(5, 1, 11, 6) = tiles 5-10, rows 1-5)
//   - Pixel-based games: tileSize=1, bounds in pixels (e.g., Rect(100, 50, 200, 150) = pixel area)
//   - UI elements: tileSize=1, bounds in pixels for buttons/widgets
type InteractionManager struct {
	objects       []InteractiveObject
	tileSize      int                                  // Grid size in pixels (1 for pixel-perfect, 16 for 16x16 tiles, etc.)
	onInteraction func(name string, gridX, gridY int) // Callback for any interaction (gridX/gridY in grid coordinates)
}

// NewInteractionManager creates a new interaction manager.
// The tileSize parameter defines the grid size in pixels:
//   - Use 1 for pixel-perfect coordinates (non-tile-based games)
//   - Use 16 for 16x16 tile grids (tile-based games)
//   - Use any value that matches your game's coordinate system
func NewInteractionManager(tileSize int) *InteractionManager {
	return &InteractionManager{
		tileSize: tileSize,
	}
}

// AddObject adds an interactive object to the manager.
func (im *InteractionManager) AddObject(obj InteractiveObject) {
	im.objects = append(im.objects, obj)
}

// SetOnInteraction sets a callback that fires for any interaction.
// The callback receives the object name and grid coordinates (not pixel coordinates).
func (im *InteractionManager) SetOnInteraction(callback func(name string, gridX, gridY int)) {
	im.onInteraction = callback
}

// Update checks for clicks and triggers callbacks.
// Converts mouse pixel coordinates to grid coordinates and checks for intersections.
func (im *InteractionManager) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		gridX := mouseX / im.tileSize
		gridY := mouseY / im.tileSize

		// Check each object
		for _, obj := range im.objects {
			if obj.Bounds.Min.X <= gridX && gridX < obj.Bounds.Max.X &&
				obj.Bounds.Min.Y <= gridY && gridY < obj.Bounds.Max.Y {

				// Call global interaction callback if set
				if im.onInteraction != nil {
					im.onInteraction(obj.Name, gridX, gridY)
				}

				// Call object-specific callback if set
				if obj.OnClick != nil {
					obj.OnClick()
				}
				return
			}
		}
	}
}
