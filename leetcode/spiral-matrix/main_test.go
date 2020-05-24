package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want []int
}{
	{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := spiralOrder(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
