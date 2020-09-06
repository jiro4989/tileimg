package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type drawParam struct {
	min, max image.Rectangle

	strokeColor color.RGBA
	lineWidth   float64
}

func draw(dest *image.RGBA, p drawParam) {
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.lineWidth)

	gc.BeginPath()
	gc.MoveTo(float64(p.min.Min.X), float64(p.min.Min.Y))
	gc.LineTo(float64(p.max.Max.X), float64(p.min.Min.Y))
	gc.LineTo(float64(p.max.Max.X), float64(p.max.Max.Y))
	gc.LineTo(float64(p.min.Min.X), float64(p.max.Max.Y))
	gc.Close()
	gc.FillStroke()
}

func drawBackground(dest *image.RGBA, c color.RGBA) {
	bounds := dest.Bounds().Max
	width := bounds.X
	height := bounds.Y
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dest.Set(x, y, c)
		}
	}
}
