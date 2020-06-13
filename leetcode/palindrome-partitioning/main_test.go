package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want [][]string
}{
	{"aab", [][]string{[]string{"a", "a", "b"}, []string{"aa", "b"}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := partition(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
