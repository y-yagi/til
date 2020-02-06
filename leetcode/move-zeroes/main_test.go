package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want []int
}{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{0, 0, 1}, []int{1, 0, 0}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		moveZeroes(tt.in)
		if !reflect.DeepEqual(tt.in, tt.want) {
			t.Fatalf("in %v, want %v", tt.in, tt.want)
		}
	}
}
