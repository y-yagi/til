package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want [][]int
}{
	{[][]int{[]int{1, 4}, []int{5, 6}}, [][]int{[]int{1, 4}, []int{5, 6}}},
	{[][]int{[]int{1, 4}, []int{2, 3}}, [][]int{[]int{1, 4}}},
	{[][]int{[]int{1, 4}, []int{1, 4}}, [][]int{[]int{1, 4}}},
	{[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}, [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}}},
	{[][]int{[]int{1, 4}, []int{4, 5}}, [][]int{[]int{1, 5}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := merge(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
