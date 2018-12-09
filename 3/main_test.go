package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fabricCoord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []coord
	}{
		{
			name: "small square",
			s:    "#6 @ 1,1: 1x1",
			want: []coord{{x: 1, y: 1}},
		},
		{
			name: "starts at zero",
			s:    "#6 @ 0,0: 1x1",
			want: []coord{{x: 0, y: 0}},
		},
		{
			name: "zero size",
			s:    "#6 @ 0,0: 0x1",
			want: nil,
		},
		{
			name: "not a square",
			s:    "#6 @ 0,1: 1x2",
			want: []coord{{x: 0, y: 1}, {x: 0, y: 2}},
		},
		{
			name: "not a square2",
			s:    "#6 @ 1,2: 1x2",
			want: []coord{{x: 1, y: 2}, {x: 1, y: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fabricCoord(tt.s))
		})
	}
}
