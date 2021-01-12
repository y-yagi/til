package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want string
}{
	{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
	{[]string{"w", "wo", "wor", "worl", "world"}, "world"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := longestWord(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
