package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  [][]int
	want bool
}{
	{2, [][]int{[]int{1, 0}}, true},
	{2, [][]int{[]int{1, 0}, []int{0, 1}}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := canFinish(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
