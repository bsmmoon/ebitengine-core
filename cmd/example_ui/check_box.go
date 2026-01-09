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

package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	checkboxWidth       = 16
	checkboxHeight      = 16
	checkboxPaddingLeft = 8
)

type CheckBox struct {
	X    int
	Y    int
	Text string

	checked   bool
	mouseDown bool

	onCheckChanged func(c *CheckBox)
}

func (c *CheckBox) width(ctx *GameContext) int {
	w := text.Advance(c.Text, &text.GoTextFace{
		Source: ctx.uiFaceSource,
		Size:   uiFontSize,
	})
	return checkboxWidth + checkboxPaddingLeft + int(w)
}

func (c *CheckBox) Update(ctx *GameContext) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if c.X <= x && x < c.X+c.width(ctx) && c.Y <= y && y < c.Y+checkboxHeight {
			c.mouseDown = true
		} else {
			c.mouseDown = false
		}
	} else {
		if c.mouseDown {
			c.checked = !c.checked
			if c.onCheckChanged != nil {
				c.onCheckChanged(c)
			}
		}
		c.mouseDown = false
	}
}

func (c *CheckBox) Draw(dst *ebiten.Image, ctx *GameContext) {
	t := imageTypeCheckBox
	if c.mouseDown {
		t = imageTypeCheckBoxPressed
	}
	r := image.Rect(c.X, c.Y, c.X+checkboxWidth, c.Y+checkboxHeight)
	ctx.drawNinePatches(dst, r, t)
	if c.checked {
		ctx.drawNinePatches(dst, r, imageTypeCheckBoxMark)
	}

	x := c.X + checkboxWidth + checkboxPaddingLeft
	y := c.Y + checkboxHeight/2
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color.Black)
	op.LineSpacing = lineSpacingInPixels
	op.PrimaryAlign = text.AlignStart
	op.SecondaryAlign = text.AlignCenter
	text.Draw(dst, c.Text, &text.GoTextFace{
		Source: ctx.uiFaceSource,
		Size:   uiFontSize,
	}, op)
}

func (c *CheckBox) Checked() bool {
	return c.checked
}

func (c *CheckBox) SetOnCheckChanged(f func(c *CheckBox)) {
	c.onCheckChanged = f
}
