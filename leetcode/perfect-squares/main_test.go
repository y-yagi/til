package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want int
}{
	{12, 3},
	{13, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := numSquares(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
