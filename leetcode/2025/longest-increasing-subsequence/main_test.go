package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{[]int{0, 1, 0, 3, 2, 3}, 4},
	{[]int{7, 7, 7, 7, 7}, 1},
}

func TestLength(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLIS(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
