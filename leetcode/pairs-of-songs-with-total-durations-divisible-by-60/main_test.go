package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{30, 20, 150, 100, 40}, 3},
	{[]int{60, 60, 60}, 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := numPairsDivisibleBy60(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
