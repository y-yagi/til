package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  int
	want []string
}{
	{[]string{"aaa", "aa", "a"}, 1, []string{"a"}},
	{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := topKFrequent(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
