package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{-1, 3, -5}, []int{-5, -1, 3}},
	{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := sortArray(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
