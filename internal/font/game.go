package font

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

// Game implements ebiten.Game interface for the font demo.
type Game struct {
	screenWidth    int
	screenHeight   int
	fontSource     *text.GoTextFaceSource
	counter        int
	kanjiText      string
	kanjiTextColor color.RGBA
}

// NewGame creates a new font game with the given configuration.
func NewGame(cfg GameConfig) *Game {
	// Load font
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		screenWidth:  cfg.ScreenWidth,
		screenHeight: cfg.ScreenHeight,
		fontSource:   fontSource,
	}
}

// Update implements ebiten.Game interface.
func (g *Game) Update() error {
	// Change the text color for each second.
	if g.counter%ebiten.TPS() == 0 {
		g.kanjiText = ""
		for j := 0; j < 6; j++ {
			for i := 0; i < 12; i++ {
				g.kanjiText += string(jaKanjis[rand.IntN(len(jaKanjis))])
			}
			g.kanjiText += "\n"
		}

		g.kanjiTextColor.R = 0x80 + uint8(rand.IntN(0x7f))
		g.kanjiTextColor.G = 0x80 + uint8(rand.IntN(0x7f))
		g.kanjiTextColor.B = 0x80 + uint8(rand.IntN(0x7f))
		g.kanjiTextColor.A = 0xff
	}
	g.counter++
	return nil
}

// Draw implements ebiten.Game interface.
func (g *Game) Draw(screen *ebiten.Image) {
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)

	const x = 20

	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 20)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: g.fontSource,
		Size:   normalFontSize,
	}, op)

	// Draw the sample text
	op = &text.DrawOptions{}
	op.GeoM.Translate(x, 60)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: g.fontSource,
		Size:   normalFontSize,
	}, op)

	// Draw Kanji text lines
	op = &text.DrawOptions{}
	op.GeoM.Translate(x, 110)
	op.ColorScale.ScaleWithColor(g.kanjiTextColor)
	op.LineSpacing = bigFontSize * 1.2
	text.Draw(screen, g.kanjiText, &text.GoTextFace{
		Source: g.fontSource,
		Size:   bigFontSize,
	}, op)
}

// Layout implements ebiten.Game interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
