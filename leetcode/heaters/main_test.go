package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	want int
}{
	{[]int{1, 2, 3}, []int{2}, 1},
	{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findRadius(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
