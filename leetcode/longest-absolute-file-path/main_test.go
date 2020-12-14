package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
	{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := lengthLongestPath(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
