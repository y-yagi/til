package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  []int
	want int
}{
	{7, []int{2, 3, 1, 2, 4, 3}, 2},
	{4, []int{1, 4, 4}, 1},
	{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	{15, []int{5, 1, 3, 5, 10, 7}, 2},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := minSubArrayLen(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
