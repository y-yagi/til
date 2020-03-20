package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  float64
	in2  int
	want float64
}{
	{2.0000, 10, 1024.00000},
	{2.10000, 3, 9.261000000000001},
	{2.00000, -2, 0.25000},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := myPow(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
