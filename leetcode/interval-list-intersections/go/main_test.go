package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  [][]int
	want [][]int
}{

	{[][]int{[]int{0, 2}, []int{5, 10}, []int{13, 23}, []int{24, 25}}, [][]int{[]int{1, 5}, []int{8, 12}, []int{15, 24}, []int{25, 26}}, [][]int{[]int{1, 2}, []int{5, 5}, []int{8, 10}, []int{15, 23}, []int{24, 24}, []int{25, 25}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := intervalIntersection(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
