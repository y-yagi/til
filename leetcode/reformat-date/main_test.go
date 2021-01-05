package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"20th Oct 2052", "2052-10-20"},
	{"26th May 1960", "1960-05-26"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := reformatDate(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
