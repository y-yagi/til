package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  []int
	want int
}{
	{15, []int{1, 2, 3, 4, 5}, 5},
	{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 8},
	{7, []int{2, 3, 1, 2, 4, 3}, 2},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := minSubArrayLen(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
