package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{1, 5}, []int{3, 4}}, 3},
	{[][]int{[]int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}}, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxDistance(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
