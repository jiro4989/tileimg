package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

type param struct {
	x, y, column, row, pad int

	strokeColor color.RGBA
	lineWidth   float64
}

func draw(dest *image.RGBA, p param) {
	gc := draw2dimg.NewGraphicContext(dest)

	// rp := rectParam{
	// 	x: p.x,
	// 	y: p.y,
	// 	column: p.column,
	// 	row: p.row,
	// }
	// r := rectangle(rp)

	gc.SetStrokeColor(p.strokeColor)
	gc.SetLineWidth(p.lineWidth)
	gc.Close()
}
