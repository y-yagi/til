package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2,
	}, {
		[][]int{{0, 1}, {0, 0}}, 1,
	},
	{
		[][]int{{1, 0}}, 0,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := uniquePathsWithObstacles(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
