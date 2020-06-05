package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]byte
	want [][]byte
}{
	{[][]byte{[]byte{'X', 'X', 'X', 'X'}, []byte{'X', 'O', 'O', 'X'}, []byte{'X', 'X', 'O', 'X'}, []byte{'X', 'O', 'X', 'X'}},
		[][]byte{[]byte{'X', 'X', 'X', 'X'}, []byte{'X', 'X', 'X', 'X'}, []byte{'X', 'X', 'X', 'X'}, []byte{'X', 'O', 'X', 'X'}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		solve(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("in '%v', want %v", tt.in1, tt.want)
		}
	}
}
