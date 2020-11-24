package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	counter := Constructor()

	// Hit at timestamp 1.
	counter.Hit(1)

	// Hit at timestamp 2.
	counter.Hit(2)

	// Hit at timestamp 3.
	counter.Hit(3)

	// get Hits at timestamp 4, should return 3.
	got := counter.GetHits(4)
	want := 3
	if got != want {
		t.Fatalf("got '%v', want '%v'", got, want)
	}

	// Hit at timestamp 300.
	counter.Hit(300)

	// get Hits at timestamp 300, should return 4.
	counter.GetHits(300)

	// get Hits at timestamp 301, should return 3.
	counter.GetHits(301)
}
