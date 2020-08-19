package main

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

var tests = []struct {
	in1  int
	in2  int
	want []int
}{
	{3, 7, []int{181, 292, 707, 818, 929}},
	{2, 1, []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := numsSameConsecDiff(tt.in1, tt.in2)
		if !reflectext.IgnoreOrderEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
