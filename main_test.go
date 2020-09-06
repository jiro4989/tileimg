package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxXY(t *testing.T) {
	tests := []struct {
		desc    string
		s       string
		wantX   int
		wantY   int
		wantX2  int
		wantY2  int
		wantErr bool
	}{
		{
			desc:    "normal: 1,2",
			s:       "1,2",
			wantX:   1,
			wantY:   2,
			wantX2:  1,
			wantY2:  2,
			wantErr: false,
		},
		{
			desc:    "normal: 1-2,2",
			s:       "1-2,2",
			wantX:   1,
			wantY:   2,
			wantX2:  2,
			wantY2:  2,
			wantErr: false,
		},
		{
			desc:    "normal: 1,2-3",
			s:       "1,2-3",
			wantX:   1,
			wantY:   2,
			wantX2:  1,
			wantY2:  3,
			wantErr: false,
		},
		{
			desc:    "normal: 1-2,3-4",
			s:       "1-2,3-4",
			wantX:   1,
			wantY:   3,
			wantX2:  2,
			wantY2:  4,
			wantErr: false,
		},
		{
			desc:    "illegal: not integer",
			s:       "foobar",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "1,foobar",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "foobar,1",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "a-2,3-4",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "1-a,3-4",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "1-2,a-4",
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "1-2,3-b",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			gotX, gotY, gotX2, gotY2, err := minMaxXY(tt.s)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.Equal(tt.wantX, gotX)
			a.Equal(tt.wantY, gotY)
			a.Equal(tt.wantX2, gotX2)
			a.Equal(tt.wantY2, gotY2)
		})
	}
}

func TestSplitHyphen(t *testing.T) {
	tests := []struct {
		desc    string
		s       string
		wantA   int
		wantB   int
		wantErr bool
	}{
		{
			desc:    "normal: 2 value",
			s:       "1-2",
			wantA:   1,
			wantB:   2,
			wantErr: false,
		},
		{
			desc:    "normal: 1 value",
			s:       "3",
			wantA:   3,
			wantB:   3,
			wantErr: false,
		},
		{
			desc:    "illegal: not integer",
			s:       "foobar",
			wantA:   0,
			wantB:   0,
			wantErr: true,
		},
		{
			desc:    "illegal: not integer",
			s:       "0-foobar",
			wantA:   0,
			wantB:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			gotA, gotB, err := splitHyphen(tt.s)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.Equal(tt.wantA, gotA)
			a.Equal(tt.wantB, gotB)
		})
	}

}
