package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{5, 4, -1, 7, 8}, 23},
	{[]int{-2, -1}, -1},
	{[]int{-2, -3, -1}, -1},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := maxSubArray(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
