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
	{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	{100, 300, []int{123, 234}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := sequentialDigits(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
