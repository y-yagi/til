package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want string
}{
	{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
	{[]string{"z", "x"}, "zx"},
	{[]string{"z", "x", "z"}, ""},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := alienOrder(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
