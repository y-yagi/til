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
	{4, 2, [][]int{[]int{2, 4}, []int{3, 4}, []int{2, 3}, []int{1, 2}, []int{1, 3}, []int{1, 4}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := combine(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
