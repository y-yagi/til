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
	{[][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}, [][]int{[]int{1, 5}, []int{6, 9}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := insert(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
