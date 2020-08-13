package main

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

var tests = []struct {
	in1  string
	want []string
}{
	{"()())()", []string{"()()()", "(())()"}},
	{"(a)())()", []string{"(a)()()", "(a())()"}},
	{")(", []string{""}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := removeInvalidParentheses(tt.in1)
		if !reflectext.IgnoreOrderEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
