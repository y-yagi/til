package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want int
}{
	{38, 2},
	{18, 9},
	{0, 0},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range tests {
		got := addDigits(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
