package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{100}, 1},
	{[]int{3, 0, 6, 1, 5}, 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := hIndex(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
