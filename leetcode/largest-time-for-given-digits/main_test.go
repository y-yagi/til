package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want string
}{
	{[]int{1, 2, 3, 4}, "23:41"},
	{[]int{5, 5, 5, 5}, ""},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := largestTimeFromDigits(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
