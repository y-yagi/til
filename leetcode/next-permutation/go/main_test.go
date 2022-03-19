package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{1, 3, 2}, []int{2, 1, 3}},
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{1, 1, 5}, []int{1, 5, 1}},
	{[]int{1, 2}, []int{2, 1}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		nextPermutation(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("in1 %v, want %v", tt.in1, tt.want)
		}
	}
}
