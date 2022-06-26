package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want []int
}{
	{0, []int{1}},
	{3, []int{1, 3, 3, 1}},
}

func TestGetRow(t *testing.T) {
	for _, tt := range tests {
		got := getRow(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
