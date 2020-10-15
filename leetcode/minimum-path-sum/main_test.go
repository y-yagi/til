package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{1, 2}, []int{1, 1}}, 3},
	{[][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}, 7},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minPathSum(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
