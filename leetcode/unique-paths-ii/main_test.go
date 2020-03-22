package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{1, 0}}, 0},
	{[][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 2},
	{[][]int{[]int{1}}, 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := uniquePathsWithObstacles(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
