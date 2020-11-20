package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)

	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)

	got := undergroundSystem.GetAverageTime("Paradise", "Cambridge") // return 14.00000. There was only one travel from "Paradise" (at time 8) to "Cambridge" (at time 22)
	if got != 14.0 {
		t.Fatalf("got '%v', want '%v'", got, 14.0)
	}
	got = undergroundSystem.GetAverageTime("Leyton", "Waterloo") // return 11.00000. There were two travels from "Leyton" to "Waterloo", a customer with id=45 from time=3 to time=15 and a customer with id=27 from time=10 to time=20. So the average time is ( (15-3) + (20-10) ) / 2 = 11.00000
	if got != 11.0 {
		t.Fatalf("got '%v', want '%v'", got, 11.0)
	}

	undergroundSystem.CheckIn(10, "Leyton", 24)
	undergroundSystem.GetAverageTime("Leyton", "Waterloo") // return 11.00000

	undergroundSystem.CheckOut(10, "Waterloo", 38)
	undergroundSystem.GetAverageTime("Leyton", "Waterloo")
}
