package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want [][]int
}{
	{1, [][]int{[]int{1}}},
	{3, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}}},
	{5, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}},
}

func TestGenerate(t *testing.T) {
	for _, tt := range tests {
		got := generate(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
