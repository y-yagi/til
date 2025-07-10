package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   string
	want int
}{
	{"42", 42},
	{"-91283472332", -2147483648},
	{"   -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := myAtoi(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
