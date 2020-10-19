package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	want int
}{
	{[]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}, 2},
	{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minDominoRotations(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
