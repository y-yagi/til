package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  int
	in3  int
	in4  int
	want [][]int
}{
	{[][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2, [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := floodFill(tt.in1, tt.in2, tt.in3, tt.in4)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
