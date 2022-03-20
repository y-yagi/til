package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want [][]string
}{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{[]string{"eat", "tea", "ate"}, []string{"tan", "nat"}, []string{"bat"}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := groupAnagrams(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
