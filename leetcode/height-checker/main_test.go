package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{1, 1, 4, 2, 1, 3}, 3},
	{[]int{5, 1, 2, 3, 4}, 5},
	{[]int{1, 2, 3, 4, 5}, 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := heightChecker(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
