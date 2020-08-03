package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{-1, -1, -1, 0, 1, 1}, 0},
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{1, 2, 3}, -1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := pivotIndex(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
