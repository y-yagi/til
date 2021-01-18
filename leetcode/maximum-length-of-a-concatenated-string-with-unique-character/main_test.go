package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"un", "iq", "ue"}, 4},
	{[]string{"cha", "r", "act", "ers"}, 6},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maxLength(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
