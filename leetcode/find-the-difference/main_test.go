package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want byte
}{
	{"abcd", "abcde", 'e'},
	{"", "y", 'y'},
	{"a", "aa", 'a'},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := findTheDifference(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1, '%v', got, '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
