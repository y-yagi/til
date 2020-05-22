package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []byte
	in2  int
	want int
}{
	{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := leastInterval(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
