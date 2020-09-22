package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  int
	want bool
}{
	{[][]int{[]int{2, 1, 5}, []int{3, 3, 7}}, 4, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := carPooling(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
