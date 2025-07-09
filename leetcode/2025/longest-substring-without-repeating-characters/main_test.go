package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   string
	want int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"", 0},
	{" ", 1},
	{"dvdf", 3},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
