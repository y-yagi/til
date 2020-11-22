package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	kv := Constructor()
	kv.Set("foo", "bar", 1) // store the key "foo" and value "bar" along with timestamp = 1
	kv.Get("foo", 1)        // output "bar"
	kv.Get("foo", 3)        // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
	kv.Set("foo", "bar2", 4)
	kv.Get("foo", 4)        // output "bar2"
	got := kv.Get("foo", 5) //output "bar2"
	if got != "bar2" {
		t.Fatalf("got '%v', want '%v'", got, "bar2")
	}
}

func TestSuccess2(t *testing.T) {
	kv := Constructor()
	kv.Set("love", "high", 10)
	kv.Set("love", "low", 20)
	got := kv.Get("love", 5)
	if got != "" {
		t.Fatalf("got '%v', want '%v'", got, "")
	}
	kv.Get("love", 10)
	got = kv.Get("love", 15)
	if got != "high" {
		t.Fatalf("got '%v', want '%v'", got, "high")
	}
	kv.Get("love", 20)
	kv.Get("love", 25)
}
