package main

import "testing"

var tests = []struct {
	in   string
	want bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isPalindrome(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
