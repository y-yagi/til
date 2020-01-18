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
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{-1}, 2, []int{-1}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		rotate(tt.in1, tt.in2)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("in1 %v, want %v", tt.in1, tt.want)
		}
	}
}
