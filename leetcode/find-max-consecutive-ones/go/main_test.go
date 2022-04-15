package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{1, 0, 1, 0, 1, 1, 1, 1}, 6},
	{[]int{1, 1, 0, 1}, 4},
	{[]int{1, 0, 1, 1, 0}, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findMaxConsecutiveOnes(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
