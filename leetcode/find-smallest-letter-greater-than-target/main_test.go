package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []byte
	in2  byte
	want byte
}{
	{[]byte{'c', 'f', 'j'}, 'j', 'c'},
	{[]byte{'c', 'f', 'j'}, 'c', 'f'},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := nextGreatestLetter(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
