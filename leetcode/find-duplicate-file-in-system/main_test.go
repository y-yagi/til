package main

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

var tests = []struct {
	in1  []string
	want [][]string
}{
	{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}, [][]string{[]string{"root/a/1.txt", "root/c/3.txt"}, []string{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findDuplicate(tt.in1)
		if !reflectext.IgnoreOrderEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
