package main

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {
	tests := []struct {
		desc  string
		param rectangleParam
		want  image.Rectangle
	}{
		{
			desc: "normal: 0, 0",
			param: rectangleParam{
				x:      0,
				y:      0,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(5, 5, 95, 395),
		},
		{
			desc: "normal: 1, 0",
			param: rectangleParam{
				x:      1,
				y:      0,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(105, 5, 195, 395),
		},
		{
			desc: "normal: 0, 1",
			param: rectangleParam{
				x:      0,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(5, 405, 95, 795),
		},
		{
			desc: "normal: 1, 1",
			param: rectangleParam{
				x:      1,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(105, 405, 195, 795),
		},
		{
			desc: "normal: 3, 1",
			param: rectangleParam{
				x:      3,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(305, 405, 395, 795),
		},
		{
			desc: "illegal: 4, 2 (out of range)",
			param: rectangleParam{
				x:      4,
				y:      2,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(305, 405, 395, 795),
		},
		{
			desc: "illegal: -1, -1 (out of range)",
			param: rectangleParam{
				x:      -1,
				y:      -1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: image.Rect(5, 5, 95, 395),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			got := rectangle(tt.param)
			a.Equal(tt.want, got)
		})
	}

}
