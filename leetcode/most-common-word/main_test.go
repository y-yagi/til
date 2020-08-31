package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  []string
	want string
}{
	{"Bob", []string{""}, "bob"},
	{"a, a, a, a, b,b,b,c, c", []string{"a"}, "b"},
	{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := mostCommonWord(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
