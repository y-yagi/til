package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := orangesRotting(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
