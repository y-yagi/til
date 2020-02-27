package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []byte
	want int
}{
	{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
	{[]byte{'a'}, 1},
	{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	{[]byte{'a', 'a', 'a', 'b', 'b'}, 4},
	{[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}, 6},
	{[]byte{'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}, 8},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := compress(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %s, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
