package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{2, 2, 2, 2, 2}, 2},
	{[]int{1, 3, 4, 2, 2}, 2},
	{[]int{3, 1, 3, 4, 2}, 3},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findDuplicate(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
