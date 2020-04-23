package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := maxArea(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
