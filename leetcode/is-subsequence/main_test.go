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
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"acb", "ahbgdc", false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isSubsequence(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', in2 '%v', got '%v', want '%v'", tt.in1, tt.in2, got, tt.want)
		}
	}
}
