package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want [][]int
}{
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[][]int{
			[]int{7, 4, 1},
			[]int{8, 5, 2},
			[]int{9, 6, 3},
		},
	},
	{
		[][]int{
			[]int{5, 1, 9, 11},
			[]int{2, 4, 8, 10},
			[]int{13, 3, 6, 7},
			[]int{15, 14, 12, 16},
		},
		[][]int{
			[]int{15, 13, 2, 5},
			[]int{14, 3, 4, 1},
			[]int{12, 6, 8, 9},
			[]int{16, 7, 10, 11},
		},
	},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		rotate(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("got %v, want %v", tt.in1, tt.want)
		}
	}
}
