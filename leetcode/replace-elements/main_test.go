package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want []int
}{
	{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := replaceElements(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
