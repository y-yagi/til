package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 10, 5, 25, 2, 8}, 28},
	{[]int{8, 10, 2}, 10},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findMaximumXOR(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
