package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want int
}{
	{1, 1},
	{5, 2},
	{9, 3},
	{15, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := consecutiveNumbersSum(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
