package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"172.16.254.1", "IPv4"},
	{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
	{"256.256.256.256", "Neither"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := validIPAddress(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
