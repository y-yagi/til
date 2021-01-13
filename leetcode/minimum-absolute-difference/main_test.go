package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{[]int{4, 2, 1, 3}, [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}}},
	{[]int{1, 3, 6, 10, 15}, [][]int{[]int{1, 3}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minimumAbsDifference(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
