package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"bcabc", "abc"},
	{"cbacdcbc", "acdb"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := removeDuplicateLetters(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
