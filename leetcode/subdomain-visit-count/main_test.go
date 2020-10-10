package main

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

var tests = []struct {
	in1  []string
	want []string
}{
	{[]string{"9001 discuss.leetcode.com"}, []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}},
	{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := subdomainVisits(tt.in1)
		if !reflectext.IgnoreOrderEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
