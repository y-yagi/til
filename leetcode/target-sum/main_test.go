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
	{[]int{1, 1, 1, 1, 1}, 3, 5},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findTargetSumWays(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
