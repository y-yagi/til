package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want bool
}{
	{[][]int{[]int{1}, []int{2}, []int{3}, []int{}}, true},
	{[][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := canVisitAllRooms(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
