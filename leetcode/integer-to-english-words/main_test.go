package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want string
}{
	{123, "One Hundred Twenty Three"},
	{12345, "Twelve Thousand Three Hundred Forty Five"},
	{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	{1234567891, "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := numberToWords(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
