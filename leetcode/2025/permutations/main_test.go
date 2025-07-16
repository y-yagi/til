package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{in1: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	{in1: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}}},
	{in1: []int{1}, want: [][]int{{1}}},
}

func TestCombinationSum(t *testing.T) {
	for _, tt := range tests {
		got := permute(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
