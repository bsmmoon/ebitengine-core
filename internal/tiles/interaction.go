package tiles

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InteractiveObject represents a clickable area in the game.
type InteractiveObject struct {
	Name    string
	Bounds  image.Rectangle // In tile coordinates
	OnClick func()
}

// InteractionManager handles click detection for interactive objects.
type InteractionManager struct {
	objects       []InteractiveObject
	tileSize      int
	onInteraction func(name string, tileX, tileY int) // Callback for any interaction
}

// NewInteractionManager creates a new interaction manager.
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
func (im *InteractionManager) SetOnInteraction(callback func(name string, tileX, tileY int)) {
	im.onInteraction = callback
}

// Update checks for clicks and triggers callbacks.
func (im *InteractionManager) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		tileX := mouseX / im.tileSize
		tileY := mouseY / im.tileSize

		// Check each object
		for _, obj := range im.objects {
			if obj.Bounds.Min.X <= tileX && tileX < obj.Bounds.Max.X &&
				obj.Bounds.Min.Y <= tileY && tileY < obj.Bounds.Max.Y {
				
				// Call global interaction callback if set
				if im.onInteraction != nil {
					im.onInteraction(obj.Name, tileX, tileY)
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

