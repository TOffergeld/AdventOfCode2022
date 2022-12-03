package main

import "testing"

func TestSumPossessions(t *testing.T) {
	want := 23
	is := sumPossessions([]int{8, 3, 6, 6})
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}

func TestFindMax(t *testing.T) {
	want := 23
	elfs := make([]Elf, 5)
	for i, v := range []int{23, 19, -200, 0, 3} {
		elfs[i].Possessions = v
	}
	is := findMax(elfs)
	if want != is {
		t.Errorf("Wanted %d, got %d", want, is)
	}
}
