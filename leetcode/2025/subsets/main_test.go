package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	{[]int{0}, [][]int{{}, {0}}},
	{[]int{}, [][]int{{}}},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := subsets(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
