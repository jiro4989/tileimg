package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
