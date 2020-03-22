package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{[]int{1, 2, 3}, [][]int{[]int{3}, []int{1}, []int{2}, []int{1, 2, 3}, []int{1, 3}, []int{2, 3}, []int{1, 2}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := subsets(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
