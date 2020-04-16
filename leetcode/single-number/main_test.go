package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{2, 2, 1}, 1},
	{[]int{4, 1, 2, 1, 2}, 4},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := singleNumber(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
