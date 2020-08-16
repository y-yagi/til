package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want []int
}{
	{"DDIIDI", []int{3, 2, 1, 4, 6, 5, 7}},
	{"DDIIII", []int{3, 2, 1, 4, 5, 6, 7}},
	{"I", []int{1, 2}},
	{"DI", []int{2, 1, 3}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findPermutation(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
