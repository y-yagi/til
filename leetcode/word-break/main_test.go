package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  []string
	want bool
}{
	{"cars", []string{"car", "ca", "rs"}, true},
	{"leetcode", []string{"leet", "code"}, true},
	{"applepenapple", []string{"apple", "pen"}, true},
	{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := wordBreak(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
