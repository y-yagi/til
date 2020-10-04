package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{2, 0, 1}, []int{2, 0, 0}},
	{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := countSmaller(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
