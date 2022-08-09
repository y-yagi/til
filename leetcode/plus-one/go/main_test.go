package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range tests {
		got := plusOne(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
