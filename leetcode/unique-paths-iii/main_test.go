package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{1, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 2, -1}}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := uniquePathsIII(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
