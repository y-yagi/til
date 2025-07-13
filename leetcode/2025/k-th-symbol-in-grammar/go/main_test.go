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
	{1, 1, 0},
	{2, 1, 0},
	{2, 2, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := kthGrammar(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
