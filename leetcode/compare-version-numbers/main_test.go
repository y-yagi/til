package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want int
}{
	{"1.0", "1", 0},
	{"1.2", "1.10", -1},
	{"1.01", "1.001", 0},
	{"0.1", "1.1", -1},
	{"1.0.1", "1", 1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := compareVersion(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
