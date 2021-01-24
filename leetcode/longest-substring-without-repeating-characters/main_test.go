package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{" ", 1},
	{"abba", 2},
	{"pwwkew", 3},
	{"abcabcbb", 3},
	{"bbbbb", 1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
