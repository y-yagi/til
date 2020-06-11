package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{0}, []int{1}, []int{5}}, 3},
	{[][]int{[]int{1, 2}}, 2},
	{[][]int{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}, 4},
	{[][]int{[]int{3, 4, 5}, []int{3, 2, 6}, []int{2, 2, 1}}, 4},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := longestIncreasingPath(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got '%v', in '%v', want %v", got, tt.in1, tt.want)
		}
	}
}
