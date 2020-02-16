package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"leetcode", 0},
	{"loveleetcode", 2},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := firstUniqChar(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1, '%v', got, '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
