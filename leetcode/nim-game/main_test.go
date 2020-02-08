package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{1, true},
	{4, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := canWinNim(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
