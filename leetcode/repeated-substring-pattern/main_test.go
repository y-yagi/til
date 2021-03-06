package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want bool
}{
	{"ababab", true},
	{"abac", false},
	{"abab", true},
	{"aba", false},
	{"abcabcabcabc", true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := repeatedSubstringPattern(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
