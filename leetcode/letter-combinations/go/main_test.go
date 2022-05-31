package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want []string
}{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := letterCombinations(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
