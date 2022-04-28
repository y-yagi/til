package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]byte
	want int
}{
	{[][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}, 1},
	{[][]byte{[]byte{'1', '1', '0', '0', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '1', '0', '0'}, []byte{'0', '0', '0', '1', '1'}}, 3},
	{[][]byte{[]byte{'1', '1', '1'}, []byte{'0', '1', '0'}, []byte{'1', '1', '1'}}, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := numIslands(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
