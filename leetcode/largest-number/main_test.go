package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want string
}{
	{[]int{10, 2}, "210"},
	{[]int{3, 30, 34, 5, 9}, "9534330"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := largestNumber(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
