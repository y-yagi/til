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
	{[]int{3, 2, 1, 5, 6, 4}, 2, 5}, {[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findKthLargest(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
