package main

import (
	"bytes"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

func main() {
	// Decode the embedded PNG
	img, err := png.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}

	// Save to file
	f, err := os.Create("tiles_spritesheet.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	log.Println("Spritesheet saved to tiles_spritesheet.png")
	log.Printf("Image size: %dx%d pixels", img.Bounds().Dx(), img.Bounds().Dy())
	log.Printf("Tiles per row: %d", img.Bounds().Dx()/16)
	log.Printf("Total rows: %d", img.Bounds().Dy()/16)
}
