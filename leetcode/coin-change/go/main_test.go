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
	{[]int{1, 2, 5}, 11, 3},
	{[]int{186, 419, 83, 408}, 6249, 20},
	{[]int{2}, 3, -1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := coinChange(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
