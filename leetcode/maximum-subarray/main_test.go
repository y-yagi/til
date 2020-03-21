package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxSubArray(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
