package helloworld

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// TestLayout ensures the game reports its intended screen size.
func TestLayout(t *testing.T) {
	g := &Game{}
	w, h := g.Layout(800, 600)
	if w != 320 || h != 240 {
		t.Errorf("expected (320,240), got (%d,%d)", w, h)
	}
}

// TestUpdate ensures Update returns no error.
func TestUpdate(t *testing.T) {
	g := &Game{}
	if err := g.Update(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

// TestDraw runs Draw with an offscreen image to ensure no panic.
func TestDraw(t *testing.T) {
	g := &Game{}
	img := ebiten.NewImage(320, 240)
	// If Draw panics, the test will fail.
	g.Draw(img)
}
