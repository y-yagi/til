package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want []string
}{
	{"AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
	{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findRepeatedDnaSequences(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
