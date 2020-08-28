package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want []int
}{
	{[][]int{[]int{4, 5}, []int{2, 3}, []int{1, 2}}, []int{-1, 0, 1}},
	{[][]int{[]int{1, 2}}, []int{-1}},
	{[][]int{[]int{3, 4}, []int{2, 3}, []int{1, 2}}, []int{-1, 0, 1}},
	{[][]int{[]int{1, 4}, []int{2, 3}, []int{3, 4}}, []int{-1, 2, -1}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findRightInterval(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
