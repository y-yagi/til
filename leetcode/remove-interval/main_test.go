package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  []int
	want [][]int
}{
	{[][]int{[]int{0, 2}, []int{3, 4}, []int{5, 7}}, []int{1, 6}, [][]int{[]int{0, 1}, []int{6, 7}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := removeInterval(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
