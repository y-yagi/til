package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"1+1+1", 3},
	{"0*0", 0},
	{"0-2147483647", -2147483647},
	{"42+1", 43},
	{"42", 42},
	{" 3+5 / 2 ", 5},
	{"3+2*2", 7},
	{" 3/2 ", 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := calculate(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
