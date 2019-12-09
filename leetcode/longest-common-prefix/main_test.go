package main

import "testing"

var tests = []struct {
	in  []string
	out string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, tt := range tests {
		got := longestCommonPrefix(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
