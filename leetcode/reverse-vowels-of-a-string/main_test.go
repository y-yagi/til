package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   string
	want string
}{
	{"hello", "holle"},
	{"leetcode", "leotcede"},
	{"aA", "Aa"},
	{"race car", "race car"},
	{"Euston saw I was not Sue.", "euston saw I was not SuE."},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := reverseVowels(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got, '%v', want '%v'", tt.in, got, tt.want)
		}
	}
}
