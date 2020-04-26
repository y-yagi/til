package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want string
}{
	{"acbbaca", "aba", "baca"},
	{"a", "aa", ""},
	{"a", "b", ""},
	{"ADOBECODEBANC", "ABC", "BANC"},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := minWindow(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
