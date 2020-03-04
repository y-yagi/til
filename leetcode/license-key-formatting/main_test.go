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
	{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
	{"2-5g-3-J", 2, "2-5G-3J"},
	{"--a-a-a-a--", 2, "AA-AA"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := licenseKeyFormatting(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}

func BenchmarkLicenseKeyFormatting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		licenseKeyFormatting("5F3Z-2e-9-w", 4)
	}
}

func BenchmarkLicenseKeyFormattingOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		licenseKeyFormatting_old("5F3Z-2e-9-w", 4)
	}
}
