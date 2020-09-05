package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainNormal(t *testing.T) {
	tests := []struct {
		desc  string
		param rectParam
		want  rect
	}{
		{
			desc: "normal: 0, 0",
			param: rectParam{
				x:      0,
				y:      0,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      5,
				y:      5,
				width:  90,
				height: 390,
			},
		},
		{
			desc: "normal: 1, 0",
			param: rectParam{
				x:      1,
				y:      0,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      105,
				y:      5,
				width:  90,
				height: 390,
			},
		},
		{
			desc: "normal: 0, 1",
			param: rectParam{
				x:      0,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      5,
				y:      405,
				width:  90,
				height: 390,
			},
		},
		{
			desc: "normal: 1, 1",
			param: rectParam{
				x:      1,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      105,
				y:      405,
				width:  90,
				height: 390,
			},
		},
		{
			desc: "normal: 3, 1",
			param: rectParam{
				x:      3,
				y:      1,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      305,
				y:      405,
				width:  90,
				height: 390,
			},
		},
		{
			desc: "illegal: 4, 2 (out of range)",
			param: rectParam{
				x:      4,
				y:      2,
				width:  400,
				height: 800,
				column: 4,
				row:    2,
				pad:    5,
			},
			want: rect{
				x:      305,
				y:      405,
				width:  90,
				height: 390,
			},
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
