package example_ui

import (
	"image"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestDrawNinePatches(t *testing.T) {
	src := ebiten.NewImage(4, 4)
	src.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	dst := ebiten.NewImage(8, 8)
	srcRect := image.Rect(0, 0, 4, 4)
	dstRect := image.Rect(0, 0, 8, 8)

	// The test passes if this does not panic.
	DrawNinePatches(src, dst, dstRect, srcRect)
}
