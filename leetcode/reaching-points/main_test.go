package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	in3  int
	in4  int
	want bool
}{
	{1, 1, 3, 5, true},
	{1, 1, 2, 2, false},
	{1, 1, 1, 1, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := reachingPoints(tt.in1, tt.in2, tt.in3, tt.in4)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
