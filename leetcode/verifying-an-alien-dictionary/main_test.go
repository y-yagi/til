package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  string
	want bool
}{
	{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
	{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := isAlienSorted(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
