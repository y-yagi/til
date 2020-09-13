package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want [][]int
}{
	{3, 9, [][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}}},
	{3, 7, [][]int{[]int{1, 2, 4}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := combinationSum3(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
