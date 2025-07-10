package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want bool
}{
	{"acb", "ahbgdc", false},
	{"abc", "ahbgdc", true},
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"", "ahbgdc", true},
	{"abc", "", false},
	{"abc", "abc", true},
	{"a", "a", true},
	{"a", "b", false},
	{"a", "aa", true},
	{"aa", "a", false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isSubsequence(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
