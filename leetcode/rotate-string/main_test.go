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
	{"abcde", "cdeab", true},
	{"abcde", "abced", false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := rotateString(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
