package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  int
	want [][]int
}{
	{[][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2, [][]int{[]int{3, 3}, []int{-2, 4}}},
	{[][]int{[]int{1, 3}, []int{-2, 2}}, 1, [][]int{[]int{-2, 2}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := kClosest(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
