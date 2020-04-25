package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{2, 3, 1, 1, 4}, true},
	{[]int{3, 2, 1, 0, 4}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := canJump(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
