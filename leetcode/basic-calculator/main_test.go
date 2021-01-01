package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"1 + 1", 2},
	{"(1+(4+5+2)-3)+(6+8)", 23},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := calculate(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
