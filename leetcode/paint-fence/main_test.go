package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want int
}{
	{2, 3, 9},
	{0, 1, 0},
	{2, 1, 1},
	{3, 2, 6},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := numWays(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
