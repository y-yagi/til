package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{1, 3, 5, 4, 7}, 2},
	{[]int{2, 2, 2, 2, 2}, 5},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findNumberOfLIS(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
