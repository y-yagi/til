package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want int
}{
	{[]int{100}, 50, 0},
	{[]int{100, 200}, 150, 1},
	{[]int{100, 200, 300, 400}, 200, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := bagOfTokensScore(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
