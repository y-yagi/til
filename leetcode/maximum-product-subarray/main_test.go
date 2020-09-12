package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{-3, 0, 1, -2}, 1},
	{[]int{0, 2}, 2},
	{[]int{2, 3, -2, 4}, 6},
	{[]int{-2, 0, -1}, 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxProduct(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
