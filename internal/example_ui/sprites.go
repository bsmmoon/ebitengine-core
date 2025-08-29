package example_ui

import "image"

type ImageType int

const (
	ImageTypeButton ImageType = iota
	ImageTypeButtonPressed
	ImageTypeTextBox
	ImageTypeVScrollBarBack
	ImageTypeVScrollBarFront
	ImageTypeCheckBox
	ImageTypeCheckBoxPressed
	ImageTypeCheckBoxMark
)

var ImageSrcRects = map[ImageType]image.Rectangle{
	ImageTypeButton:          image.Rect(0, 0, 16, 16),
	ImageTypeButtonPressed:   image.Rect(16, 0, 32, 16),
	ImageTypeTextBox:         image.Rect(0, 16, 16, 32),
	ImageTypeVScrollBarBack:  image.Rect(16, 16, 24, 32),
	ImageTypeVScrollBarFront: image.Rect(24, 16, 32, 32),
	ImageTypeCheckBox:        image.Rect(0, 32, 16, 48),
	ImageTypeCheckBoxPressed: image.Rect(16, 32, 32, 48),
	ImageTypeCheckBoxMark:    image.Rect(32, 32, 48, 48),
}
