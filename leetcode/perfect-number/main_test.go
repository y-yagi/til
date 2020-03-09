package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want bool
}{
	{28, true},
	{6, true},
	{5, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := checkPerfectNumber(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
