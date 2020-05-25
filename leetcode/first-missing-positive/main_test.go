package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{1, 2, 0}, 3},
	{[]int{3, 4, -1, 1}, 2},
	{[]int{7, 8, 9, 11, 12}, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := firstMissingPositive(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
