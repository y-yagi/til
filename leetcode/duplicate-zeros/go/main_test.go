package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		duplicateZeros(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("in '%v', want %v", tt.in1, tt.want)
		}
	}
}
