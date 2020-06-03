package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want [][]int
}{
	{[][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{1, 0, 1}, []int{0, 1, 1}, []int{0, 1, 0}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		gameOfLife(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("in '%v', want %v", tt.in1, tt.want)
		}
	}
}
