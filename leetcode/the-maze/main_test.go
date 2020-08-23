package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  []int
	in3  []int
	want bool
}{
	{[][]int{[]int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0}, []int{1, 1, 0, 1, 1}, []int{0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4}, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := hasPath(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
