// Copyright 2017 The Ebiten Authors
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

package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	Rect image.Rectangle
	Text string

	mouseDown bool

	onPressed func(b *Button)
}

func (b *Button) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *Button) Draw(dst *ebiten.Image, ctx *GameContext) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	ctx.drawNinePatches(dst, b.Rect, t)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(b.Rect.Min.X+b.Rect.Max.X)/2, float64(b.Rect.Min.Y+b.Rect.Max.Y)/2)
	op.ColorScale.ScaleWithColor(color.Black)
	op.LineSpacing = ctx.lineSpacingInPixels
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	text.Draw(dst, b.Text, &text.GoTextFace{
		Source: ctx.uiFaceSource,
		Size:   ctx.uiFontSize,
	}, op)
}

func (b *Button) SetOnPressed(f func(b *Button)) {
	b.onPressed = f
}
