package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{4, 10, 4, 3, 8, 9}, 3},
	{[]int{10, 9, 2, 5, 3, 4}, 3},
	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLIS(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
