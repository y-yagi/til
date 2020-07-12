package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want [][]string
}{
	{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, [][]string{[]string{"a", "z"}, []string{"abc", "bcd", "xyz"}, []string{"acef"}, []string{"az", "ba"}}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := groupStrings(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
