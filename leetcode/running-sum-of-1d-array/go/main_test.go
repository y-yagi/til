package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := runningSum(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
