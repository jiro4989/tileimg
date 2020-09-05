package main

func draw(width, height, column, row, pad int) {
	// img := image.NewRGBA(image.Rect(0, 0, width, height))
}

type rectParam struct {
	x, y, column, row, width, height, pad int
}

type rect struct {
	x, y, width, height int
}

func rectangle(p rectParam) rect {
	var (
		col = p.width / p.column
		row = p.height / p.row
		xx  = col * p.x
		yy  = row * p.y
		w   = col - p.pad*2
		h   = row - p.pad*2
	)

	r := rect{
		x:      xx + p.pad,
		y:      yy + p.pad,
		width:  w,
		height: h,
	}
	return r
}
