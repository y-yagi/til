package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  int
	want string
}{
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := convert(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
