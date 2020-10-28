package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  string
	in3  string
	want int
}{
	{[]string{"a", "a", "b", "b"}, "a", "b", 1},
	{[]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice", 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := shortestDistance(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
