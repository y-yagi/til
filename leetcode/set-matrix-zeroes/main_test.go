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
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		[][]int{
			[]int{1, 0, 1},
			[]int{0, 0, 0},
			[]int{1, 0, 1},
		},
	},
	{
		[][]int{
			[]int{0, 1, 2, 0},
			[]int{3, 4, 5, 2},
			[]int{1, 3, 1, 5},
		},
		[][]int{
			[]int{0, 0, 0, 0},
			[]int{0, 4, 5, 0},
			[]int{0, 3, 1, 0},
		},
	},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		setZeroes(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("got %v, want %v", tt.in1, tt.want)
		}
	}
}
