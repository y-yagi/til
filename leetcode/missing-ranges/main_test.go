package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	in3  int
	want []string
}{
	{[]int{0}, 0, 9, []string{"1->9"}},
	{[]int{-1}, -2, 0, []string{"-2", "0"}},
	{[]int{-1}, -2, -1, []string{"-2"}},
	{[]int{-1}, -1, -1, []string{}},
	{[]int{}, 1, 1, []string{"1"}},
	{[]int{0, 1, 3, 50, 75}, 0, 99, []string{"2", "4->49", "51->74", "76->99"}},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		got := findMissingRanges(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
