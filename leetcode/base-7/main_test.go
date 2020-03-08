package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want string
}{
	{100, "202"},
	{-7, "-10"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := convertToBase7(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
