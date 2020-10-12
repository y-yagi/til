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
	{"ab", "ba", true},
	{"ab", "ab", false},
	{"aa", "aa", true},
	{"aaaaaaabc", "aaaaaaacb", true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := buddyStrings(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
