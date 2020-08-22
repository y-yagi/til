package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want string
}{
	{3, "III"},
	{9, "IX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := intToRoman(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
