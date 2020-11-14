package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"aab", "aba"},
	{"aaab", ""},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := reorganizeString(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
