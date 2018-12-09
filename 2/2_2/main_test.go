package main

import "testing"

func Test_compare(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want int
	}{
		{
			name: "sample1",
			s1: "abcde",
			s2: "bbcde",
			want: 1,
		},
		{
			name: "sample 2",
			s1: "aacdf",
			s2: "bbcde",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.s1, tt.s2); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
