package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want int
}{
	{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := subarraysDivByK(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
