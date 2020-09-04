package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want []string
}{
	{[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}, []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := reorderLogFiles(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
