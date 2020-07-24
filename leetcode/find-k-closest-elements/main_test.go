package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	in3  int
	want []int
}{

	{[]int{1, 1, 1, 10, 10, 10}, 1, 9, []int{10}},
	{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := findClosestElements(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
