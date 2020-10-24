package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{-2, 1, 2, -2, 1, 2}, true},
	{[]int{1, 0, 1, -4, -3}, false},
	{[]int{3, 5, 0, 3, 4}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{3, 1, 4, 2}, true},
	{[]int{-1, 3, 2, 0}, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := find132pattern(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
