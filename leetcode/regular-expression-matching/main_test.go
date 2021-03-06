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
	{"aa", "a", false},
	{"aa", "a*", true},
	{"aab", "c*a*b", true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := isMatch(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
