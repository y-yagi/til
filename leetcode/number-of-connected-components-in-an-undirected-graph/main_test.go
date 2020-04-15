package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  [][]int
	want int
}{
	{4, [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 2}}, 2},
	{5, [][]int{[]int{0, 1}, []int{1, 2}, []int{3, 4}}, 2},
	{5, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{3, 4}}, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := countComponents(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v' - '%v', got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
