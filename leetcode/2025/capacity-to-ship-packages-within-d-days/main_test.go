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
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
	{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
	{[]int{1, 2, 3, 1, 1}, 4, 3},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := shipWithinDays(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
