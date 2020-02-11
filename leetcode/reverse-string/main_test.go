package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   []byte
	want []byte
}{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		reverseString(tt.in)
		if !reflect.DeepEqual(tt.in, tt.want) {
			t.Fatalf("got %v, want %v", tt.in, tt.want)
		}
	}
}
