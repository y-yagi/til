package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	in3  int
	want float64
}{
	{1, 1, 1, 0.0},
	{2, 1, 1, 0.5},
	{100000009, 33, 17, 1.0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := champagneTower(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
