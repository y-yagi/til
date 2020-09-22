package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{3, 2, 3}, []int{3}},
	{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := majorityElement(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
