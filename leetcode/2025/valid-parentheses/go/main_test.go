package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isValid(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
