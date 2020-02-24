package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want int
}{
	{"Hello, my name is John", 5},
	{"Hello", 1},
	{"", 0},
	{"                ", 0},
	{", , , ,        a, eaefa", 6},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := countSegments(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
