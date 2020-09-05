package main

import (
	"image/color"
)

var (
	colorBlack = color.RGBA{0, 0, 0, 255}
	colorWhite = color.RGBA{255, 255, 255, 255}
	colors     = map[string]color.RGBA{
		"white": colorWhite,
		"black": colorBlack,
	}
)
