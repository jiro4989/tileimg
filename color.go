package main

import (
	"image/color"
)

var (
	colorNone  = color.RGBA{0, 0, 0, 0}
	colorBlack = color.RGBA{0, 0, 0, 255}
	colorRed   = color.RGBA{255, 0, 0, 255}
	colorGreen = color.RGBA{0, 255, 0, 255}
	colorBlue  = color.RGBA{0, 0, 255, 255}
	colorWhite = color.RGBA{255, 255, 255, 255}
	colors     = map[string]color.RGBA{
		"none":  colorNone,
		"black": colorBlack,
		"red":   colorRed,
		"green": colorGreen,
		"blue":  colorBlue,
		"white": colorWhite,
	}
)
