package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := twoSum(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
