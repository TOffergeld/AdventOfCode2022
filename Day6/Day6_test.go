package Day6

import "testing"

func TestSumPossessions(t *testing.T) {
	want := 5
	is := findMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}

func TestSumPossessions2(t *testing.T) {
	want := 10
	is := findMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}
