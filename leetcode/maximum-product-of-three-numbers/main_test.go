package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{-100, -2, -3, 1}, 300},
	{[]int{-100, -98, -1, 2, 3, 4}, 39200},
	{[]int{1, 2, 3}, 6},
	{[]int{1, 2, 3, 4}, 24},
	{[]int{-1, -2, -3}, -6},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maximumProduct(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
