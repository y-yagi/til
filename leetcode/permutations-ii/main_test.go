package main

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{[]int{2, 2, 1, 1}, [][]int{[]int{1, 1, 2, 2}, []int{1, 2, 1, 2}, []int{1, 2, 2, 1}, []int{2, 1, 1, 2}, []int{2, 1, 2, 1}, []int{2, 2, 1, 1}}},
	{[]int{1, 1, 2}, [][]int{[]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := permuteUnique(tt.in1)
		if !reflectext.IgnoreOrderEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
