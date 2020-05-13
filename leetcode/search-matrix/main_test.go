package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  int
	want bool
}{
	{[][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 5, true},
	{[][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 20, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := searchMatrix(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
