package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findDisappearedNumbers(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
