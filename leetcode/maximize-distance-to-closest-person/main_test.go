package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{0, 0, 0, 0, 1, 0, 1}, 4},
	{[]int{0, 1, 1, 1, 0, 0, 1, 0, 0, 0}, 3},
	{[]int{0, 1, 0, 0, 0, 0}, 4},
	{[]int{0, 0, 1, 0, 1, 1}, 2},
	{[]int{1, 0, 0, 0, 1, 0, 1}, 2},
	{[]int{1, 0, 0, 0}, 3},
	{[]int{0, 1}, 1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxDistToClosest(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
