package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{3, 7, 2, 3}, true},
	{[]int{5, 3, 4, 5}, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := stoneGame(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
