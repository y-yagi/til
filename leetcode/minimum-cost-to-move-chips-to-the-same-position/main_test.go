package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{1, 2, 3}, 1},
	{[]int{2, 2, 2, 3, 3}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minCostToMoveChips(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
