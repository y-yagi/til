package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	iterator := Constructor("abc", 2)

	if iterator.Next() != "ab" {
		t.Fatalf("Fail at first Next")
	}

	if !iterator.HasNext() {
		t.Fatalf("Fail at first HasNext")
	}

	if iterator.Next() != "ac" {
		t.Fatalf("Fail at second Next")
	}

	if !iterator.HasNext() {
		t.Fatalf("Fail at second HasNext")
	}

	if iterator.Next() != "bc" {
		t.Fatalf("Fail at third Next")
	}

	if iterator.HasNext() {
		t.Fatalf("Fail at third HasNext")
	}
}
