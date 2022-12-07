package day6

import "testing"

func TestSumPossessions(t *testing.T) {
	want := 5
	is := findPacketMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}

func TestSumPossessions2(t *testing.T) {
	want := 10
	is := findPacketMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}
