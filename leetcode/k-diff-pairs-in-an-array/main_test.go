package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want int
}{
	{[]int{3, 1, 4, 1, 5}, 2, 2},
	{[]int{1, 2, 3, 4, 5}, 1, 4},
	{[]int{1, 3, 1, 5, 4}, 0, 1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findPairs(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
