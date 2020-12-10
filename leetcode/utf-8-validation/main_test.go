package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{197, 130, 1}, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := validUtf8(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
