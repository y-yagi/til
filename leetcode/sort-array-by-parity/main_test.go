package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []int
	want []int
}{
	{[]int{0, 1}, []int{0, 1}},
	{[]int{0, 2}, []int{0, 2}},
	{[]int{3, 1, 2, 4}, []int{4, 2, 1, 3}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := sortArrayByParity(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
