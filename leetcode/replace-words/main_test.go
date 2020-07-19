package main

import (
	"testing"
)

var tests = []struct {
	in1  []string
	in2  string
	want string
}{
	{[]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa", "a a a a a a a a bbb baba a"},
	{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := replaceWords(tt.in1, tt.in2)
		if got != tt.want {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
