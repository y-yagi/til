package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want [][]string
}{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{"", ""}, [][]string{{"", ""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := groupAnagrams(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
