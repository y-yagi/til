package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{0, 1, 2, 2}, 3},
	{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
	{[]int{1, 2, 1}, 3},
	{[]int{1, 2, 3, 2, 2}, 4},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := totalFruit(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
