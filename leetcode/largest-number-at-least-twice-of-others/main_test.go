package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 6, 1, 0}, 1},
	{[]int{0, 0, 3, 2}, -1},
	{[]int{1, 2, 3, 4}, -1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := dominantIndex(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
