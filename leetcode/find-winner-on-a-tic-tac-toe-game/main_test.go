package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want string
}{
	{[][]int{[]int{0, 0}, []int{2, 0}, []int{1, 1}, []int{2, 1}, []int{2, 2}}, "A"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := tictactoe(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
