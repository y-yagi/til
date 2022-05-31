package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]byte
	in2  string
	want bool
}{
	{[][]byte{[]byte{'C', 'A', 'A'}, []byte{'A', 'A', 'A'}, []byte{'B', 'C', 'D'}}, "AAB", true},
	{[][]byte{[]byte{'a', 'a'}}, "aaa", false},
	{[][]byte{[]byte{'a', 'b'}}, "ba", true},
	{[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCCED", true},
	{[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "SEE", true},
	{[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCB", false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := exist(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
