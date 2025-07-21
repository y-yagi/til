package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]byte
	want int
}{
	{
		[][]byte{{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}},
		3,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := numIslands(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
