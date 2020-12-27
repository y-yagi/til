package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  int
	want []string
}{
	{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{"This    is    an", "example  of text", "justification.  "}},
	{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16, []string{"What   must   be", "acknowledgment  ", "shall be        "}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := fullJustify(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
