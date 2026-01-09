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
)

type Game struct {
	uiContext    *GameContext
	screenWidth  int
	screenHeight int
	button1      *Button
	button2      *Button
	checkBox     *CheckBox
	textBoxLog   *TextBox
}

func NewGame(cfg GameConfig) *Game {
	g := &Game{
		uiContext:    NewGameContext(cfg.UIFontSize, cfg.LineSpacingInPixels),
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
	}
	g.button1 = &Button{
		Rect: image.Rect(16, 16, 144, 48),
		Text: "Button 1",
	}
	g.button2 = &Button{
		Rect: image.Rect(160, 16, 288, 48),
		Text: "Button 2",
	}
	g.checkBox = &CheckBox{
		X:    16,
		Y:    64,
		Text: "Check Box!",
	}
	g.textBoxLog = &TextBox{
		Rect: image.Rect(16, 96, 624, 464),
	}

	g.button1.SetOnPressed(func(b *Button) {
		g.textBoxLog.AppendLine("Button 1 Pressed")
	})
	g.button2.SetOnPressed(func(b *Button) {
		g.textBoxLog.AppendLine("Button 2 Pressed")
	})
	g.checkBox.SetOnCheckChanged(func(c *CheckBox) {
		msg := "Check box check changed"
		if c.Checked() {
			msg += " (Checked)"
		} else {
			msg += " (Unchecked)"
		}
		g.textBoxLog.AppendLine(msg)
	})
	return g
}

func (g *Game) Update() error {
	g.button1.Update()
	g.button2.Update()
	g.checkBox.Update(g.uiContext)
	g.textBoxLog.Update(g.uiContext)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xeb, 0xeb, 0xeb, 0xff})
	g.button1.Draw(screen, g.uiContext)
	g.button2.Draw(screen, g.uiContext)
	g.checkBox.Draw(screen, g.uiContext)
	g.textBoxLog.Draw(screen, g.uiContext)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
