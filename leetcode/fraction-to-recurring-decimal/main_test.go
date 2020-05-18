package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want string
}{
	{4, 333, "0.(012)"},
	{6, 3, "2"},
	{1, 2, "0.5"},
	{2, 1, "2"},
	{2, 3, "0.(6)"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := fractionToDecimal(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
