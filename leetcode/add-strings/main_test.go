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
	{"1", "2", "3"},
	{"5", "7", "12"},
	{"120", "3921", "4041"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := addStrings(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
