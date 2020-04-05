package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	in3  int
	want [][]int
}{
	{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{[]int{1, 2}, []int{1, 4}, []int{1, 6}}},
	{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{[]int{1, 1}, []int{1, 1}}},
	{[]int{1, 2}, []int{3}, 3, [][]int{[]int{1, 3}, []int{2, 3}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := kSmallestPairs(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
