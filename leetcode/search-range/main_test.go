package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want []int
}{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{2, 2}, 2, []int{0, 1}},
	{[]int{1}, 1, []int{0, 0}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := searchRange(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
