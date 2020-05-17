package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want int
}{
	{10, 3, 3},
	{7, -3, -2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := divide(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
