package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want bool
}{
	{2, true},
	{3, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := divisorGame(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
