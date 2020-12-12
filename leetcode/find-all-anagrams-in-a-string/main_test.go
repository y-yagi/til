package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want []int
}{
	{"cbaebabacd", "abc", []int{0, 6}},
	{"abab", "ab", []int{0, 1, 2}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findAnagrams(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
