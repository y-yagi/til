package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want [][]int
}{
	{[]int{2, 3, 6, 7}, 7, [][]int{[]int{2, 2, 3}, []int{7}}},
	{[]int{2, 3, 5}, 8, [][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := combinationSum(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
