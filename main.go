package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

const usage = `layoutimg is TODO

Usage:
  layoutimg [options]
  layoutimg -h | --help
  layoutimg --version

Options:
  -h, --help
      --version
  -W, --width <width>
  -H, --height <height>
  -c, --column <column>
  -r, --row <row>
  -b, --background-color <background-color>
  -l, --line-color <line-color>
  -f, --fill-color <fill-color>

`

func main() {
	dest := image.NewRGBA(image.Rect(0, 0, 400, 200))
	drawBackground(dest, color.RGBA{255, 255, 255, 255})

	dp := drawParam{
		x:           0,
		y:           0,
		column:      4,
		row:         2,
		pad:         5,
		strokeColor: color.RGBA{0, 0, 0, 255},
		lineWidth:   3,
	}
	draw(dest, dp)

	dp = drawParam{
		x:           1,
		y:           0,
		column:      4,
		row:         2,
		pad:         5,
		strokeColor: color.RGBA{0, 0, 0, 255},
		lineWidth:   3,
	}
	draw(dest, dp)

	draw2dimg.SaveToPngFile("out.png", dest)
}
