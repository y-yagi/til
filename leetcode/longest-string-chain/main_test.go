package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"a", "b", "ab", "bac"}, 2},
	{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
	{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
	{[]string{"a", "b", "ba", "abc", "abd", "bdca"}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := longestStrChain(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
