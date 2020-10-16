package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"12", 2},
	{"226", 3},
	{"0", 0},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := numDecodings(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
