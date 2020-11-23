package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  int
	want string
}{
	{"deeedbbcccbdaa", 3, "aa"},
	{"abcd", 2, "abcd"},
	{"pbbcggttciiippooaais", 2, "ps"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := removeDuplicates(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
