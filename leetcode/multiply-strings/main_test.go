package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want string
}{
	{"2", "3", "6"},
	{"123", "456", "56088"},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := multiply(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
