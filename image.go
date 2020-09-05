package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type drawParam struct {
	x, y, column, row, pad int

	strokeColor color.RGBA
	lineWidth   float64
}

func draw(dest *image.RGBA, p drawParam) {
	gc := draw2dimg.NewGraphicContext(dest)

	rp := rectParam{
		x:      p.x,
		y:      p.y,
		column: p.column,
		row:    p.row,
	}
	r := rectangle(rp)

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.lineWidth)

	// left top
	gc.MoveTo(float64(r.Min.X), float64(r.Min.Y))
	gc.LineTo(float64(r.Max.X-r.Min.X), float64(r.Min.Y))
	gc.LineTo(float64(r.Min.X), float64(r.Max.Y-r.Min.Y))

	// right bottom
	gc.MoveTo(float64(r.Max.X), float64(r.Max.Y))
	gc.LineTo(float64(r.Min.X-r.Max.X), float64(r.Max.Y))
	gc.LineTo(float64(r.Max.X), float64(r.Min.Y-r.Max.Y))

	gc.Close()
	gc.FillStroke()
}
