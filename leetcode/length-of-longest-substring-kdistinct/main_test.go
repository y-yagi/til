package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  int
	want int
}{
	{"a", 0, 0},
	{"aa", 1, 2},
	{"eceba", 2, 3},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLongestSubstringKDistinct(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
