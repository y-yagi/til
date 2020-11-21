package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want bool
}{
	{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := search(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
