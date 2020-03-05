package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{1, 1, 0, 1, 1, 1}, 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findMaxConsecutiveOnes(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
