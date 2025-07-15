package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want [][]int
}{
	{
		in1:  []int{2, 3, 6, 7},
		in2:  7,
		want: [][]int{{2, 2, 3}, {7}},
	},
	{
		in1:  []int{2, 3, 5},
		in2:  8,
		want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
	},
	{
		in1:  []int{2},
		in2:  1,
		want: [][]int{},
	},
}

func TestCombinationSum(t *testing.T) {
	for _, tt := range tests {
		got := combinationSum(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
