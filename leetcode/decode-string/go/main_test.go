package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"3[a2[c]]", "accaccacc"},
	{"3[a]2[bc]", "aaabcbc"},
	{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	{"abc3[cd]xyz", "abccdcdcdxyz"},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := decodeString(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
