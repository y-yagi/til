package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  []int
	want int
}{
	{4, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 15},
	{2, []int{2, 4, 1}, 2},
	{2, []int{3, 2, 6, 5, 0, 3}, 7},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxProfit(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
