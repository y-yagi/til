package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"5", "2", "C", "D", "+"}, 30},
	{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := calPoints(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
