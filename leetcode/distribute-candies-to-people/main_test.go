package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want []int
}{
	{7, 4, []int{1, 2, 3, 1}},
	{10, 3, []int{5, 2, 3}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := distributeCandies(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
