package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"cc", 2},
	{"leetcode", 2},
	{"abbcccddddeeeeedcba", 5},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxPower(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
