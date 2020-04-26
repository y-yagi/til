package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{12, 345, 2, 6, 7896}, 2},
	{[]int{555, 901, 482, 1771}, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findNumbers(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
