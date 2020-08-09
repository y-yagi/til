package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"a good   example", "example good a"},
	{"  hello world!  ", "world! hello"},
	{"the sky is blue", "blue is sky the"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := reverseWords(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
