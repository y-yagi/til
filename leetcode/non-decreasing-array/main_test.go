package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{3, 4, 2, 3}, false},
	{[]int{4, 2, 3}, true},
	{[]int{4, 2, 1}, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := checkPossibility(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
