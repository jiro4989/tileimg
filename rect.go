package main

import (
	"image"
)

type rectParam struct {
	x, y, column, row, width, height, pad int
}

func rectangle(p rectParam) image.Rectangle {
	col := p.width / p.column
	row := p.height / p.row

	if p.x < 0 {
		p.x = 0
	}

	if p.y < 0 {
		p.y = 0
	}

	if p.column-1 < p.x {
		p.x = p.column - 1
	}

	if p.row-1 < p.y {
		p.y = p.row - 1
	}

	xx := col*p.x + p.pad
	yy := row*p.y + p.pad
	w := col - p.pad*2
	h := row - p.pad*2

	r := image.Rect(xx, yy, xx+w, yy+h)
	return r
}
