package main

import (
	"log"
	"testing"
)

func TestSuccess(t *testing.T) {
	s1 := Constructor([]int{1, 0, 0, 2, 3})
	s2 := Constructor([]int{0, 3, 0, 4, 0})
	got := s1.dotProduct(s2)
	if got != 8 {
		log.Fatal(got)
	}

	s1 = Constructor([]int{9, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	s2 = Constructor([]int{6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	got = s1.dotProduct(s2)
	if got != 54 {
		log.Fatal(got)
	}
}
