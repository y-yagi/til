package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]byte
	want int
}{
	{[][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}, 4},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := maximalSquare(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
