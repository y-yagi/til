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
	{"abba", "dog cat cat dog", true},
	{"abba", "dog cat cat fish", false},
	{"aaaa", "dog cat cat dog", false},
	{"abba", "dog dog dog dog", false},
	{"ab", "dog dog", false},
	{"aaa", "aa aa aa aa", false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := wordPattern(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, in2 '%v', got, %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
