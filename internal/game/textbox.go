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
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	textBoxPaddingLeft = 8
	textBoxPaddingTop  = 4
)

type TextBox struct {
	Rect image.Rectangle
	Text string

	vScrollBar *VScrollBar
	offsetX    int
	offsetY    int
}

func (t *TextBox) AppendLine(line string) {
	if t.Text == "" {
		t.Text = line
	} else {
		t.Text += "\n" + line
	}
}

func (t *TextBox) Update(ctx *GameContext) {
	if t.vScrollBar == nil {
		t.vScrollBar = &VScrollBar{}
	}
	t.vScrollBar.X = t.Rect.Max.X - VScrollBarWidth
	t.vScrollBar.Y = t.Rect.Min.Y
	t.vScrollBar.Height = t.Rect.Dy()

	_, h := t.contentSize(ctx)
	t.vScrollBar.Update(h)

	t.offsetX = 0
	t.offsetY = t.vScrollBar.ContentOffset()
}

func (t *TextBox) contentSize(ctx *GameContext) (int, int) {
	h := int(float64(len(strings.Split(t.Text, "\n")))*ctx.lineSpacingInPixels) + textBoxPaddingTop
	return t.Rect.Dx(), h
}

func (t *TextBox) viewSize() (int, int) {
	return t.Rect.Dx() - VScrollBarWidth - textBoxPaddingLeft, t.Rect.Dy()
}

func (t *TextBox) contentOffset() (int, int) {
	return t.offsetX, t.offsetY
}

func (t *TextBox) Draw(dst *ebiten.Image, ctx *GameContext) {
	ctx.drawNinePatches(dst, t.Rect, imageTypeTextBox)

	textOp := &text.DrawOptions{}
	x := -float64(t.offsetX) + textBoxPaddingLeft
	y := -float64(t.offsetY) + textBoxPaddingTop
	textOp.GeoM.Translate(x, y)
	textOp.GeoM.Translate(float64(t.Rect.Min.X), float64(t.Rect.Min.Y))
	textOp.ColorScale.ScaleWithColor(color.Black)
	textOp.LineSpacing = ctx.lineSpacingInPixels
	text.Draw(dst.SubImage(t.Rect).(*ebiten.Image), t.Text, &text.GoTextFace{
		Source: ctx.uiFaceSource,
		Size:   ctx.uiFontSize,
	}, textOp)

	t.vScrollBar.Draw(dst, ctx)
}
