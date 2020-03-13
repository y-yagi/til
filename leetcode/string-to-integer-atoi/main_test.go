package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"42", 42},
	{"  -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
	{"-91283472332", -2147483648},
	{"+1", 1},
	{" +0 123", 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := myAtoi(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
