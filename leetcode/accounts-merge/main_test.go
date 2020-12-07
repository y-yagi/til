package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]string
	want [][]string
}{
	{
		[][]string{[]string{"John", "johnsmith@mail.com", "john00@mail.com"}, []string{"John", "johnnybravo@mail.com"}, []string{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, []string{"Mary", "mary@mail.com"}},
		[][]string{[]string{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, []string{"John", "johnnybravo@mail.com"}, []string{"Mary", "mary@mail.com"}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := accountsMerge(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
