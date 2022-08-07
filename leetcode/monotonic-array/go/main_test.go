package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{1, 1, 0}, true},
	{[]int{1, 2, 2, 3}, true},
	{[]int{6, 5, 4, 4}, true},
	{[]int{1, 3, 2}, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := isMonotonic(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
