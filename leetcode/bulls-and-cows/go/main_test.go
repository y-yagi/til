package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want string
}{
	{"1807", "7810", "1A3B"},
	{"1123", "0111", "1A1B"},
	{"1122", "2211", "0A4B"},
	{"1122", "1222", "3A0B"},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := getHint(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
