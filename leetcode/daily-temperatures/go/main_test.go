package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want []int
}{
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := dailyTemperatures(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
