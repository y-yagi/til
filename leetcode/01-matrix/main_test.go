package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want [][]int
}{
	{[][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 2, 1}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := updateMatrix(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
