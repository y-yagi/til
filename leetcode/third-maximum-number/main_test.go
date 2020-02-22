package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 2, 1}, 1},
	{[]int{1, 2}, 2},
	{[]int{2, 2, 3, 1}, 1},
	{[]int{1, 1, 2}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := thirdMax(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
