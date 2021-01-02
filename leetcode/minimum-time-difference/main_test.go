package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"23:59", "00:00"}, 1},
	{[]string{"00:00", "23:59", "00:00"}, 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findMinDifference(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
